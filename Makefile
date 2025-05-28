format:
	go fmt ./...

vet: format
	go vet ./...

build: vet
	CGO_ENABLED=0 go build ./cmd/traverser/main.go
