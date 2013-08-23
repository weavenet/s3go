all: deps test
	@mkdir -p bin/
	go build -v -o bin/s3go .
deps:
	go get -d -v ./...
test: deps
	go test ./...
