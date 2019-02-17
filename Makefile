.PHONY: all

all: \
	dist/workspace-darwin-amd64 \
	dist/workspace-linux-amd64 \
	dist/workspace-windows-amd64.exe

build:
dist/workspace-darwin-amd64:
	mkdir -p dist
	GOOS=darwin GOARCH=amd64 go build -o $@ ./src/workspace

dist/workspace-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o $@ ./src/workspace

dist/workspace-windows-amd64.exe:
	GOOS=windows GOARCH=amd64 go build -o $@ ./src/workspace

clean:
	rm -rf dist
	
