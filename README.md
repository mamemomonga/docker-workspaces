# Dockerで利用する汎用作業環境

Dockerで使う汎用の作業環境です。

git, vim, screen など、シェルでの各種作業を行う前提のいろいろと便利なものを詰め込んだものです。

コンテナ内部にappユーザが作成され、
ローカルの home/app が [docker-volume-bindfs](https://github.com/lebokus/docker-volume-bindfs) によりコンテナの /home/app にマウントされます。
これを経由することにより、コンテナ内部のappユーザ・グループをローカルのユーザ・グループにマッピングするようにしています。

* docker-volume-bindfs を導入する際に特権が求められるので許可をお願いします。

* 本イメージは開発・作業用です。プロダクション環境での利用はおすすめしません。
* /home/app 以下以外の変更はコンテナ終了後に消えます。

https://hub.docker.com/r/mamemomonga/workspaces

# イメージ一覧

ブランチ | ディストリビューション    | 環境
---------|---------------------------|----
[debian](https://github.com/mamemomonga/docker-workspaces/tree/debian) | Debian 9 (stretch)        | 汎用
[ubuntu](https://github.com/mamemomonga/docker-workspaces/tree/ubuntu) | Ubuntu 18.04 LTS (buster) | 汎用

# 使い方 

## クイックスタート

Linux

	$ mkdir ws
	$ cd ws
	$ curl -o workspace https://raw.githubusercontent.com/mamemomonga/docker-workspaces/master/dist/workspace-linux-amd64
	$ chmod 755 ./workspace

Docker for macOS

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

	$ ./workspace.sh stop

