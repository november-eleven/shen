all: dist

dist: client server lint test

client:
	scripts/dist-client

server:
	scripts/dist-server

clean:
	scripts/clean

lint: client server
	scripts/lint

test: client server
	scripts/test

archive: dist
	scripts/archive

.PHONY: client server dist lint test archive
