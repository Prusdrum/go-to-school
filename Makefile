GO_VERSION=1.21.4

install:
	mkdir -p bin

run-cli:
	go run ./cmd/cli/main.go