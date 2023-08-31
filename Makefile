.PHONY: pre-commit lint test

# Keep this rule first to make it the default
pre-commit: lint test

lint:
	golangci-lint run

test:
	go test -race -covermode atomic -v ./...
	go build -o /dev/null
