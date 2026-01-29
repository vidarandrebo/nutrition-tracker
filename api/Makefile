.PHONY: format
format:
	find . -type f -name "*.go" | xargs gofumpt -w

.PHONY: generate
generate:
	go generate ./...
	find . -type f -name "*.go" | xargs gofumpt -w
