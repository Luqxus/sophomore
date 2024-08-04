build:
	@go build -o ./bin/spaces


run: build
	@./bin/spaces

test:
	@go test ./...