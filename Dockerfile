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
	cd /tmp && unzip terraform.zip && mv terraform /usr/bin && rm -f terraform.zip
