.PHONY: format
format:
	find . -type f -name "*.go" | xargs gofumpt -w

.PHONY: openapi
openapi:
	go tool oapi-codegen -config oapi-codegen.yaml api.yaml
