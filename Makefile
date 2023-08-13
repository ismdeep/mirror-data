all: github-tasks jetbrains nodejs adoptium openssl

github-tasks:
	go run ./cmd/github-tasks

jetbrains:
	go run ./cmd/jetbrains

nodejs:
	go run ./cmd/nodejs

adoptium:
	go run ./cmd/adoptium

openssl:
	go run ./cmd/openssl
