#!/bin/bash
set -eu
BASEDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

if [ ! -e "$BASEDIR/config" ]; then
	echo "config file not exists."
	exit 1
fi

source $BASEDIR/config

do_pull() {
	docker pull $IMAGE_NAME
}

do_home() {
	if [ -d "$VOL_MOUNT" ]; then
		echo "$VOL_MOUNT already exists."
		exit 1
	fi
	mkdir -p $VOL_MOUNT
	docker run --rm $IMAGE_NAME tar cC /home/app-skel . | tar xvC $VOL_MOUNT
}

do_start() {

	if [ -n "$(docker ps -q -f name=$CONTAINER_NAME)" ]; then
		echo "$CONTAINER_NAME already exists."
		exit 1
	fi

	if [ ! -e "$VOL_MOUNT" ]; then
		echo "$VOL_MOUNT not exists."
		exit 1
	fi

	if [ -n "$(docker volume ls -q -f name=$VOL_NAME)" ]; then
		echo "$VOL_NAME already exists."
		exit 1
	fi

	# bindfsプラグインがなければインストールする
	if [ -z $(docker plugin ls --format '{{.Name}}' | grep 'lebokus/bindfs') ]; then
		echo "Install docker plugin lebokus/bindfs"
		docker plugin install lebokus/bindfs
	fi

	# bindfsプラグインが無効ならば有効にする
	if [ $(docker plugin ls --format '{{.Name}} {{.Enabled}}' | grep lebokus/bindfs:latest | awk '{print $2}') == "false" ]; then
		echo "Set Enable plugin lebokus/bindfs"
		docker plugin enable lebokus/bindfs
	fi

	#  Dockerが動作しているカーネルとローカルのカーネルが同じ
	if [ "$(docker info --format '{{.KernelVersion}}')" == "$(uname -r)" ]; then
		# おそらくLinuxである
		# UID,GIDをそれぞれ1000, ローカルユーザ・グループにマッピングした
		# ボリュームの作成
		echo "Create Volume(Linux): $VOL_NAME"
		docker volume create \
		    -d lebokus/bindfs \
		    -o sourcePath=$VOL_MOUNT \
		    -o map=$(id -u)/10000:@$(id -g)/@10000 \
			$VOL_NAME > /dev/null
	else
		# おそらくDocker for Macなどである
		# UID,GIDをそれぞれ0(root)にマッピングした
		# ボリュームの作成
		echo "Create Volume(Mac/Windows): $VOL_NAME"
		docker volume create \
		    -d lebokus/bindfs \
		    -o sourcePath=$VOL_MOUNT \
		    -o map=0/10000:@0/@10000 \
			$VOL_NAME > /dev/null
	fi

	docker run --rm -it --hostname $CONTAINER_NAME --name $CONTAINER_NAME \
		-v $VOL_NAME:/home/app \
		-d \
		$IMAGE_NAME	sleep infinity
	echo "Start Container: $CONTAINER_NAME"
}

do_stop() {
	echo "Stop Container: $CONTAINER_NAME"
	docker rm -f $CONTAINER_NAME || true
	echo "Remove Volume: $VOL_NAME"
	if [ -n "$(docker volume ls -q -f name=$VOL_NAME)" ]; then
		docker volume rm $VOL_NAME
	fi
}

usage() {
	echo "USAGE"
	echo " $0 [ build | pull | home ]"
	echo " $0 [ start | stop ]"
	echo " $0 [ root | app ]"
}

run() {
	case "${1:-}" in
		"start" )  do_start ;;
		"stop"  )  do_stop ;;
		"pull"  )  do_pull ;;
		"home"  )  do_home ;;
		"root"  )  exec docker exec -it $CONTAINER_NAME bash ;;
		"app"   )  exec docker exec -it $CONTAINER_NAME bash -c 'cd /home/app && exec gosu app bash' ;;
	esac
}

if [ -z "${1:-}" ]; then usage; exit 1; fi
for i in $@; do
	run $i
done

