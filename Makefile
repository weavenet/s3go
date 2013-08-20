all: deps
	@mkdir -p bin/
	go build -v -o bin/s3go .
deps:
	@echo "Installing Deps"
	go get -d -v ./...
test: deps
	@echo "Testing s3go"
	go test ./...
