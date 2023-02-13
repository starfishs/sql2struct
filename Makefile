test:
	CGO_ENABLED=0  go test -v ./...
golangci-lint:
	GOOS=linux golangci-lint run --timeout 5m -v
.PHONY: test
