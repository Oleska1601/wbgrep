.PHONY: test vet lint fmt build

test:
	go test -v -race ./...

vet:
	go vet ./...

lint:
	golint ./...

fmt:
	go fmt ./...
	goimports -w .

check: fmt vet lint test

build:
	go build -o sorter ./cmd/app