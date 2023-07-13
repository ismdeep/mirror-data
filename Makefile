all: install go docker-compose openssl ctop adoptium nodejs harbor alpine-linux python image-syncer

install:
	pip install -r requirements.txt

go:
	python go.py

docker-compose:
	python docker-compose.py

openssl:
	python openssl.py

ctop:
	python ctop.py

adoptium:
	python adoptium.py

nodejs:
	python nodejs.py

harbor:
	python harbor.py

alpine-linux:
	python alpine-linux.py

python:
	python python.py

image-syncer:
	python image-syncer.py
