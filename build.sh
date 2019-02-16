#!/bin/bash
set -eu
BASEDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

if [ ! -e "$BASEDIR/config" ]; then
	echo "config file not exists."
	exit 1
fi

source $BASEDIR/config

exec docker build -t $IMAGE_NAME .

