clamor:
	mkdir -p dist
	go build -C cmd/server -o ./dist/clamor

clean:
	rm -rf dist
