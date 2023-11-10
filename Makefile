run:
	./bin/mirror-data

build:
	mkdir -p ./bin/
	go build -o ./bin/mirror-data .
