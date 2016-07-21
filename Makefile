all: dist

dist: client server lint test

client:
	npm install

server:
	go build
	mv shen dist/

clean:
	rm -rf dist/*
	rm -rf node_modules/
	rm -f shen
	rm -f shen.tar.gz

lint:
	scripts/lint

test:
	scripts/test

archive: dist
	cd dist && tar -zcf '../shen.tar.gz' --exclude='.gitkeep' .

.PHONY: client server dist lint test archive
