# Dockerで利用する汎用作業環境

* Dockerで使う汎用のCLI作業環境です。
* git, vim, screen など、シェルでの各種作業を行う前提のいろいろと便利なものを詰め込んだものです。
* Dockerコマンドを実行し、プラグインの導入や起動を手助けするプログラムが付属しています(workspace)。ソースコードは[こちら](./src)。
* それぞれのDocker Imageは、ブランチごとに分けています。

コンテナ内部にappユーザが作成され、
ローカルの home/app が [docker-volume-bindfs](https://github.com/lebokus/docker-volume-bindfs) によりコンテナの /home/app にマウントされます。
これによってコンテナ内部のappユーザ・グループをローカルのユーザ・グループにマッピングするようにしていますので、
コンテナ内部とホスト側でのユーザ・グループの食い違いが発生しません。

docker-volume-bindfs導入時に以下のような質問が表示されますので、yを入力して許可してください。

	Plugin "lebokus/bindfs" is requesting the following privileges:
	 - mount: [/var/lib/docker/plugins/]
	 - mount: [/]
	 - device: [/dev/fuse]
	 - capabilities: [CAP_SYS_ADMIN]
	Do you grant the above permissions? [y/N] y

* 本イメージは開発・作業用です。プロダクション環境での利用はおすすめしません。
* /home/app 以下以外の変更は、コンテナ終了後に消えます。

# Docker Hub

https://hub.docker.com/r/mamemomonga/workspaces

# イメージ一覧

ブランチ | ディストリビューション    | 環境
---------|---------------------------|----
[debian](https://github.com/mamemomonga/docker-workspaces/tree/debian) | Debian 9 (stretch)        | 汎用
[ubuntu](https://github.com/mamemomonga/docker-workspaces/tree/ubuntu) | Ubuntu 18.04 LTS (buster) | 汎用
[cloud-infra](https://github.com/mamemomonga/docker-workspaces/tree/cloud-infra) | Debian 9 (stretch) | GCP, AWS管理用

debian

* Timezoneは Asia/Tokyoです。
* screen, openssh-client, build-essential, jqなどが導入済みです。
* 標準エディタはvimです。
* appユーザにおいてsudo コマンドはパスワードなしで実行可能です。

ubuntu

* Timezoneは Asia/Tokyoです。
* screen, openssh-client, build-essentialなどが導入済みです。
* 標準エディタはvimです。

cloud-infra

* debianに加えて terraform, Google Cloud SDK, AWSCLIが導入済みです。

# 使い方 

## クイックスタート

Docker for mac

	$ mkdir ws
	$ cd ws
	$ curl -o workspace https://raw.githubusercontent.com/mamemomonga/docker-workspaces/master/dist/workspace-darwin-amd64
	$ chmod 755 ./workspace
	$ ./workspace config-debian
	$ ./workspace pull home start
	$ ./workspace root
	$ ./workspace app
	$ ./workspace stop

## 詳細

Linux

	$ mkdir ws
	$ cd ws
	$ curl -o workspace https://raw.githubusercontent.com/mamemomonga/docker-workspaces/master/dist/workspace-linux-amd64
	$ chmod 755 ./workspace

Docker for mac

	$ mkdir ws
	$ cd ws
	$ curl -o workspace https://raw.githubusercontent.com/mamemomonga/docker-workspaces/master/dist/workspace-darwin-amd64
	$ chmod 755 ./workspace

Docker for Windows(未検証)

	> mkdir ws
	> cd ws
	> bitsadmin /TRANSFER workspace https://github.com/mamemomonga/docker-workspaces/raw/master/dist/workspace-windows-amd64.exe %CD%\workspace.exe

以下、Windowsの場合は **./workspace** を **workspace.exe** に読み替えて下さい。

初回実行

	$ ./workspace config-debian
	$ ./workspace pull home start app stop

再度起動

	$ ./workspace start app stop

## 設定テンプレートの取得

debian

	$ ./workspace config-debian

ubuntu

	$ ./workspace config-ubuntu

## 設定の確認

	$ vim config.yaml

## イメージの取得

	$ ./workspace pull

## ホームディレクトリの作成

	$ ./workspace home

## コンテナの起動

	$ ./workspace start

## rootユーザでコンテナの中にログイン

	$ ./workspace root

## appユーザでコンテナの中にログイン

	$ ./workspace app

## コンテナの停止

	$ ./workspace stop

## アンインストール

停止・イメージ削除・プラグイン削除(debianの場合)

	$ ./workspace stop
	$ docker rmi mamemomonga/workspaces:debian
	$ docker plugin disable lebokus/bindfs:latest
	$ docker plugin rm lebokus/bindfs:latest
