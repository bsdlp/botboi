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

binary:
	$(docker_go) go build -ldflags "-linkmode external -extldflags -static" -a -o build/botboi .

image:
	docker build -t bsdlp/botboi .

build: binary image

push_to_hub:
	docker push bsdlp/botboi

deploy: push_to_hub
