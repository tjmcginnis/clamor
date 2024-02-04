clamor:
	mkdir -p dist
	go build -o dist/clamor ./cmd/clamor

docker:
	docker build -t clamor .

clean:
	rm -rf dist
