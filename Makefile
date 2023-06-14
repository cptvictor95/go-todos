build:
	go build -o bin/go-todos ./cmd/go-todos

run: build
	./bin/go-todos

clean:
	rm bin/*