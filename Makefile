all: github-tasks jetbrains nodejs adoptium openssl python

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

python:
	go run ./cmd/python
