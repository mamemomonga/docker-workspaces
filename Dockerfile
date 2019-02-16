FROM debian:stretch

RUN set -xe && \
	export DEBIAN_FRONTEND=noninteractive && \
	apt-get update && \
	apt-get install -y --no-install-recommends \
		tzdata \
		locales \
		gosu \
		sudo \
		wget \
		curl \
		p7zip-full \
		ca-certificates \
		openssh-client \
		build-essential \
		apt-transport-https \
		git-core \
		vim \
		screen \
		man && \
	rm -rf /var/lib/apt/lists/*

ENV TZ Asia/Tokyo

RUN set -xe && \
	rm -f /etc/localtime && \
	ln -s /usr/share/zoneinfo/${TZ} /etc/localtime && \
	echo ${TZ} > /etc/timezone

RUN set -xe && \
	perl -i -nlpE 's!^# (en_US.UTF-8 UTF-8)!$1!; s!^# (ja_JP.UTF-8 UTF-8)!$1!; ' /etc/locale.gen && \
	locale-gen && \
	update-locale LANG=en_US.UTF-8 && \
	sh -c "echo '3' | update-alternatives --config editor"

RUN set -xe && \
	useradd -m -s /bin/bash -u 10000 app

ADD assets/ /

RUN set -xe && \
	echo 'source /etc/screenrc.local' >> /etc/screenrc

