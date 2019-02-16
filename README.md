# 作業環境

Ubuntu Bionic(18.04 LTS)の汎用作業環境です。

コンテナ内部の  の appユーザが作成され、
ローカル の home/app が bindfs により /home/app にマウントされます。

コンテナ内部の /home/app の appユーザ(UID:10000, GID:10000) のファイルは、
ローカルがLinuxの場合はローカルのGID,UIDに、
ローカルがmacOSの場合は0(root)にマッピングされ、docker for Macの機能でローカルUID,GIDにマッピングされます。

[DockerHub](https://hub.docker.com/r/mamemomonga/workspace-ubuntu-bionic)

# 使い方 

# クイックスタート

	$ mkdir workspace
	$ cd workspace
	$ curl -o workspace.sh https://raw.githubusercontent.com/mamemomonga/docker-workspaces/ubuntu/workspace.sh
	$ chmod 755 ./workspace.sh
	$ ./workspace.sh pull home start
	$ ./workspace.sh app
	$ ./workspace.sh stop

## 設定の確認

	$ ./workspace.sh
	$ vim config

## DockerHubからイメージ取得

	$ ./workspace.sh pull

## ビルド

	$ ./workspace.sh build

## ホームディレクトリの作成

	$ ./workspace.sh home

## コンテナの起動

	$ ./workspace.sh start

## rootでコンテナの中に入る

	$ ./workspace.sh root

## appでコンテナの中に入る

	$ ./workspace.sh app

## コンテナの停止

	$ ./workspace.sh stop

