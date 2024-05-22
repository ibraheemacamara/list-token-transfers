include .env
export $(shell sed 's/=.*//' .env)
build:
	@go build .

test:
	@go clean -testcache && go test ./... -cover

run:
	@go run .
