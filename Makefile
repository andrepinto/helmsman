build:
	CGO_ENABLED=0 go build -ldflags '-s' -o bin/helmsman cmd/server/main.go

dep:
	dep ensure

docker:
	docker build -f Dockerfile -t andrepinto/helmsman:1.1.0 .

.PHONY: build dep docker
