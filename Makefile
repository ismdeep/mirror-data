run:
	./bin/mirror-data

build:
	mkdir -p ./bin/
	GOOS=linux  GOARCH=amd64 go build -o ./bin/mirror-data-linux-amd64  .
	GOOS=linux  GOARCH=arm64 go build -o ./bin/mirror-data-linux-arm64  .
	GOOS=darwin GOARCH=amd64 go build -o ./bin/mirror-data-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build -o ./bin/mirror-data-darwin-arm64 .
