build:
	@go build -o bin/space-invader
run: build
	@./bin/space-invader
test:
	@go test -v ./...
