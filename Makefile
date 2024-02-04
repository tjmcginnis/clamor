clamor:
	mkdir -p dist
	go build -o dist/clamor ./cmd/clamor

clean:
	rm -rf dist
