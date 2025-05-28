.PHONY: format vet build run
TRAVERSER_PATH = ./cmd/traverser/main.go
TRAVERSER_BIN = traverser
DIST_PATH = ./dist

format:
	go fmt ./...

vet: format
	go vet ./...

run: vet
	go run $(TRAVERSER_PATH)

build: vet
	GOOS=linux GOARCH=amd64 go build -o $(DIST_PATH)/$(TRAVERSER_BIN)-linux-amd64 $(TRAVERSER_PATH)
	GOOS=darwin GOARCH=arm64 go build -o $(DIST_PATH)/$(TRAVERSER_BIN)-darwin-arm64 $(TRAVERSER_PATH)
