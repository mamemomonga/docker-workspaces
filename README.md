# Dockerで利用する汎用作業環境

Dockerで使う汎用の作業環境です。

git, vim, screen など、シェルでの各種作業を行う前提のいろいろと便利なものを詰め込んだものです。

コンテナ内部にappユーザが作成され、
ローカルの home/app が [docker-volume-bindfs](https://github.com/lebokus/docker-volume-bindfs) によりコンテナの /home/app にマウントされます。
これを経由することにより、コンテナ内部のappユーザ・グループをローカルのユーザ・グループにマッピングするようにしています。

docker-volume-bindfs を導入する際に特権が求められるので許可をお願いします。

ブランチ | ディストリビューション / 環境
---------|-------------------------------
debian   | Debian 9 (stretch)
ubuntu   | Ubuntu 18.04 LTS (buster)

本イメージは開発・作業用です。プロダクション環境での利用はおすすめしません。

https://hub.docker.com/r/mamemomonga/workspaces

# 使い方 

## クイックスタート

	$ mkdir workspace
	$ cd workspace
	$ curl -o workspace.sh https://raw.githubusercontent.com/mamemomonga/docker-workspaces/master/workspace.sh
	$ chmod 755 ./workspace.sh

debian

	$ curl -o config https://raw.githubusercontent.com/mamemomonga/docker-workspaces/debian/config

ubuntu

	$ curl -o config https://raw.githubusercontent.com/mamemomonga/docker-workspaces/ubuntu/config

以下共通

	$ ./workspace.sh pull home start app
	$ ./workspace.sh stop

再度起動してappユーザで入る

	$ ./workspace.sh start app
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

