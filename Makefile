all: dist

dist: client server

client:
	npm install

server:
	go build
	mv shen dist/

clean:
	rm -r dist/*

lint:
	scripts/lint

.PHONY: client server dist lint
