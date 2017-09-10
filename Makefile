build:
    go build -o helmsman cmd/server/main.go

dep:
    glide install

.PHONY: build dep