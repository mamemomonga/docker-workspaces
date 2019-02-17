package main

import (
	"log"
	"os"
	"fmt"
	"os/exec"
//	"github.com/davecgh/go-spew/spew"
)

func usage() {
	fmt.Println(fmt.Sprintf("Usage: %s arguments",os.Args[0]))
	fmt.Println("Arguments:")
	fmt.Println("   pull, home")
	fmt.Println("   start, stop")
	fmt.Println("   config-debian, config-ubuntu, config-cloud-infra")
	os.Exit(1)
}

var config Config

func main() {

	if len(os.Args) < 2 {
		usage()
	}
	actions := os.Args[1:]
	for _,i := range actions {
		switch i {
			case "config-debian": do_fetch_config("debian"); os.Exit(0)
			case "config-ubuntu": do_fetch_config("ubuntu"); os.Exit(0)
			case "config-cloud-infra": do_fetch_config("cloud-infra"); os.Exit(0)
		}
	}

	cfg, err := LoadConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	config = cfg

	for _,i := range actions {
		switch i {
//			case "config": do_config()
			case "home":  do_home()
			case "pull":  do_pull()
			case "start": do_start()
			case "stop":  do_stop()
			case "root":  do_root()
			case "app":   do_app()
			default: usage()
		}
	}

}

func info_run(s string) {
	log.Printf("\033[44;1m RUN \033[47;30;1m %s \033[0m",s)
}

// ホームディレクトリの作成
func do_home() {
	info_run("Create home directory")

	if _, err := os.Stat( config.Volume.Mount ); !os.IsNotExist(err) {
		log.Printf("%s already exists.\n",config.Volume.Mount)
		os.Exit(1)
	}
	if err := os.MkdirAll( config.Volume.Mount, 0755 ); err != nil {
		log.Fatal(err)
	}

	run_stdout2expand( config.Volume.Mount ,"docker","run","--rm",config.Docker.Image,"tar","cC","/home/app",".")

}

// docker pull
func do_pull() {
	info_run("Pull image")
	if err := run_command("docker", "pull", config.Docker.Image ); err != nil {
		log.Fatal(err)
	}
}

// 起動
func do_start() {
	info_run("Start container")

	// コンテナはすでに稼働していないか
	{
		b, err := exec.Command("docker","ps","-q","-f",
			fmt.Sprintf("name=%s",config.Docker.Container)).CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		if len(b) > 0 {
			log.Fatal(fmt.Sprintf("Container %s is already exists.", config.Docker.Container))
		}
	}

	// bindfsプラグインがなければインストールする
	{
		b, err := exec.Command("docker","plugin","ls","--format","{{.Name}}").CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		if string(b[0:len(b)-1]) != "lebokus/bindfs:latest" {
			log.Println("Install: docker plugin lebokus/bindfs:latest")
			if err := run_command("docker","plugin","install","lebokus/bindfs"); err != nil {
				log.Fatal(err)
			}

		}
	}

	// bindfsプラグインが無効ならば有効にする
	{
		b, err := exec.Command("docker","plugin","ls","--format","{{.Name}} {{.Enabled}}").CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		if string(b[0:len(b)-1]) != "lebokus/bindfs:latest true" {
			log.Println("Enable: docker plugin lebokus/bindfs:latest")
			if err := run_command("docker","plugin","enable","lebokus/bindfs"); err != nil {
				log.Fatal(err)
			}

		}
	}

	// Docker for Mac または Windowsの場合
	if config.Runtime.GOOS == "darwin" || config.Runtime.GOOS == "windows" {
		// UID,GIDをそれぞれ0(root)にマッピングした
		// ボリュームの作成
		log.Printf("Create Volume(%s): %s\n", config.Runtime.GOOS, config.Volume.Name)
		if err := run_command("docker","volume","create",
		    "-d","lebokus/bindfs:latest",
		    "-o", fmt.Sprintf("sourcePath=%s",config.Volume.Mount),
		    "-o","map=0/10000:@0/@10000",
			config.Volume.Name,
		); err != nil {
			log.Fatal(err)
		}
	} else if config.Runtime.GOOS == "linux" {
		// UID,GIDをそれぞれ1000, ローカルユーザ・グループにマッピングした
		// ボリュームの作成
		log.Printf("Create Volume(%s): %s\n", config.Runtime.GOOS, config.Volume.Name)
		if err := run_command("docker","volume","create",
		    "-d","lebokus/bindfs:latest",
		    "-o", fmt.Sprintf("sourcePath=%s",config.Volume.Mount),
		    "-o", fmt.Sprintf("map=%s/10000:@%s/@10000", config.Runtime.Uid, config.Runtime.Gid),
			config.Volume.Name,
		); err != nil {
			log.Fatal(err)
		}
	}

	// コンテナ起動
	log.Printf("Start Container: %s", config.Volume.Name)
	if err := run_command("docker","run","--rm","-d",
		"--hostname", config.Docker.Container,
		"--name", config.Docker.Container,
		"-v", fmt.Sprintf("%s:%s", config.Volume.Name, "/home/app"),
		config.Docker.Image,
		"sleep","infinity",
	); err != nil {
		log.Fatal(err)
	}

}

// 終了
func do_stop() {
	info_run("Stop container")
	fmt.Printf("Stop Container: %s\n", config.Docker.Container)
	run_command("docker","rm","-f", config.Docker.Container)

	fmt.Printf("Remove Volume: %s\n", config.Volume.Name)
	run_command("docker","volume","rm", config.Volume.Name)
}

// 
func do_root() {
	info_run("Login as root")
	if err := run_command("docker","exec","-it",config.Docker.Container,"login","-f","root"); err != nil {
		log.Fatal(err)
	}
}

func do_app() {
	info_run("Login as app")
	if err := run_command("docker","exec","-it",config.Docker.Container,"login","-f","app"); err != nil {
		log.Fatal(err)
	}
}

//func do_config() {
//	spew.Dump(config)
//}

