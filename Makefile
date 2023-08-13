all: github-tasks jetbrains nodejs adoptium

github-tasks:
	go run ./cmd/github-tasks

jetbrains:
	go run ./cmd/jetbrains

nodejs:
	go run ./cmd/nodejs

adoptium:
	go run ./cmd/adoptium
