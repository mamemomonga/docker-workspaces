FROM golang:1.12.1-alpine3.9

RUN set -xe && apk add make git
ADD src/ /go/src/workspace

RUN set -xe && \
	cd /go/src/workspace && \
	make deps

RUN set -xe && \
	cd /go/src/workspace && \
	make

