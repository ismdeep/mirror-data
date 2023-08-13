all: github-tasks jetbrains nodejs adoptium openssl python alpine-linux go

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

alpine-linux:
	go run ./cmd/alpine-linux

go:
	go run ./cmd/go
