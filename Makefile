go_version = 1.11
docker_workdir = /go/src/github.com/bsdlp/botboi
docker_go := docker run --rm -v $(CURDIR):$(docker_workdir) -w $(docker_workdir) golang:$(go_version)

.PHONY: deploy

update-deps:
	dep ensure -update

lint:
	$(docker_go) sh -c 'go get -u github.com/twitchtv/retool && retool do gometalinter ./...'

test:
	$(docker_go) go test ./...

build_binaries:
	$(docker_go) go build -o build/botboi .

build: build_binaries build_lambdas
