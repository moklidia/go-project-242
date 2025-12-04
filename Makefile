build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

run:
	./bin/hexlet-path-size $(path)

run-human:
	./bin/hexlet-path-size --human $(path)

lint:
	golangci-lint run

dev:
	go run ./cmd/hexlet-path-size $(path)

test:
	go test -v ./tests
