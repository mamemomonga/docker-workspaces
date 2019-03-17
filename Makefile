
DOCKER_IMAGE:=workspace_build

dist:
	mkdir -p dist
	docker build -t $(DOCKER_IMAGE) .
	# docker run --rm -it $(DOCKER_IMAGE)
	docker run --rm $(DOCKER_IMAGE) tar cC /go/src/workspace/bin . | tar xvC dist

clean:
	rm -rf dist

purge: clean
	docker rmi -f $(DOCKER_IMAGE)

.PHONY: clean purge
