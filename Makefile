all: github-tasks jetbrains

github-tasks:
	go run ./cmd/github-tasks

jetbrains:
	go run ./cmd/jetbrains
