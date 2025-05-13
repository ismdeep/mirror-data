help:
	@cat Makefile | grep '# `' | grep -v '@cat Makefile'

# `make bin`                       Build binary
.PHONY: bin
bin:
	go build -o bin/data   -mod vendor -trimpath -ldflags '-s -w' ./app/data/
	go build -o bin/mirror -mod vendor -trimpath -ldflags '-s -w' ./app/mirror/

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

# `make install-rclone`            Install rclone
.PHONY: install-rclone
install-rclone:
	mkdir -p cache/ && \
		cd cache/ && \
		curl -fLO https://github.com/rclone/rclone/releases/download/v1.69.2/rclone-v1.69.2-linux-amd64.deb
	apt install -y ./cache/rclone-v1.69.2-linux-amd64.deb

# `make test-prepare`              Prepare test
.PHONY: test-prepare
test-prepare:
	docker run --rm --name mirror-data-runtime -d golang:1.23.8-bullseye sleep infinity || true
	docker exec --workdir / mirror-data-runtime mkdir -p /root/.config/rclone/
	docker exec --workdir / mirror-data-runtime touch    /root/.config/rclone/rclone.conf

# `make test-clean`                Clean test
.PHONY: test-clean
test-clean:
	docker stop mirror-data-runtime || true
	docker rm   mirror-data-runtime || true

# `make test`                      Test
.PHONY: test
test:
	docker exec mirror-data-runtime rm -rf /mirror-data/
	docker exec mirror-data-runtime mkdir -p /mirror-data/
	docker cp ./ mirror-data-runtime:/mirror-data/
	docker exec --workdir /mirror-data/ mirror-data-runtime make install-rclone
	docker exec --workdir /mirror-data/ mirror-data-runtime make bin
	docker exec --workdir /mirror-data/ mirror-data-runtime ./bin/data

# `make clean`                     Clean
.PHONY: clean
clean:
	rm -rf bin/
