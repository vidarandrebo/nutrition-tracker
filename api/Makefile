.PHONY: format
format:
	find . -type f -name "*.go" | xargs gofumpt -w

.PHONY: openapi
openapi:
	go generate ./...
