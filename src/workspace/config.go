package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"errors"
	"os"
	"fmt"
	"runtime"
	"os/user"
	"path"
)

type Config struct {
	Docker  CDocker `yaml:"docker"`
	Volume  CVolume `yaml:"volume"`
	Runtime CRuntime
	Ports   []string `yaml:"ports"`
}

type CDocker struct {
	Image     string `yaml:image`
	Container string `yaml:container`
}

type CVolume struct {
	Name  string `yaml:name`
	Mount string `yaml:mount`
}

type CRuntime struct {
	GOOS   string
	GOARCH string
	Uid    string
	Gid    string
	Cwd    string
}

func LoadConfig(filename string) (Config, error) {

	var cnf Config

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return cnf,errors.New(fmt.Sprintf("%s not found",filename))
	}

	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return cnf, err
	}

	err = yaml.Unmarshal(buf, &cnf)
	if err != nil {
		return cnf, err
	}

	// 先頭に/があると絶対パス
	// なければカレントディレクトリからの相対パス
	cwd, err := os.Getwd()
	if err != nil {
		return cnf, err
	}
	cnf.Runtime.Cwd = cwd

	if cnf.Volume.Mount[0] != '/' {
		cnf.Volume.Mount = path.Join(cwd,cnf.Volume.Mount)
	}

	cnf.Runtime.GOOS   = runtime.GOOS
	cnf.Runtime.GOARCH = runtime.GOARCH

	user, err := user.Current()
	if err != nil {
		return cnf, err
	}
	cnf.Runtime.Uid = user.Uid
	cnf.Runtime.Gid = user.Gid

	return cnf,nil
}

