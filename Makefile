help:

all: install go docker-compose openssl ctop adoptium

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
