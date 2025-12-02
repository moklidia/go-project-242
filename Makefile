build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

run:
	./bin/hexlet-path-size $(path)

lint:
	golangci-lint run

dev:
	go run ./cmd/hexlet-path-size $(path)
