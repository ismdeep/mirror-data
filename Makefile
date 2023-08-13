all: github-tasks jetbrains nodejs

github-tasks:
	go run ./cmd/github-tasks

jetbrains:
	go run ./cmd/jetbrains

nodejs:
	go run ./cmd/nodejs
