CURRENT_BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

build:
	exec docker build -t mamemomonga/workspaces:$(CURRENT_BRANCH) .

.PHONY: build

