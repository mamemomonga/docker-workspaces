FROM mamemomonga/workspaces:debian

RUN set -xe && \
	export DEBIAN_FRONTEND=noninteractive && \
	apt-get update && \
	apt-get install -y --no-install-recommends \
		python-pip python-yaml groff && \
	rm -rf /var/lib/apt/lists/* && \
	pip install awscli

RUN set -xe && \
	curl -o /tmp/terraform.zip https://releases.hashicorp.com/terraform/0.11.11/terraform_0.11.11_linux_amd64.zip && \
	7za x -o/usr/local/bin /tmp/terraform.zip && rm -f terraform.zip

RUN set -xe && \
	curl -o /tmp/gcs.tar.gz https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-225.0.0-linux-x86_64.tar.gz && \
	tar zxpf /tmp/gcs.tar.gz -C /usr/local && \
	chown -R app:app /usr/local/google-cloud-sdk && \
	gosu app /usr/local/google-cloud-sdk/install.sh -q

