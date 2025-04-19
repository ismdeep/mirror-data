help:
	@cat Makefile | grep '# `' | grep -v '@cat Makefile'

# `make bin`                       Build binary
.PHONY: bin
bin:
	go build -o bin/data   ./app/data/
	go build -o bin/mirror ./app/mirror/

# `make publish`                   Publish
.PHONY: publish
publish:
	if [ "$$(git status --short | wc -l)" != "0" ]; then \
		git config user.name "L. Jiang" && \
		git config user.email "l.jiang.1024@gmail.com" && \
		git add . && \
		git commit -am "docs: update data via github workflow" && \
		git push origin HEAD:main && \
		docker buildx build \
			--platform linux/amd64 \
			--progress plain --pull --push -t ismdeep/mirror-server:latest -f ./app/mirror/Dockerfile .; \
	fi

# `make clean`                     Clean
.PHONY: clean
clean:
	rm -rf bin/
