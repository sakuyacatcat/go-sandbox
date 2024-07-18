BENCH ?= .
BENCHTIME ?= 5s

.PHONY: build run

build:
	echo "Building..."
	go build -o $(path)/main.exe $(path)/main.go

run:
	echo "Running..."
	go run $(path)

test:
	echo "Testing..."
	go test -v ./...

test-package:
	echo "Testing..."
	go test -v ${path}

test-coverage:
	echo "Testing..."
	go test -cover ./...

bench:
	echo "Benchmarking..."
	go test -bench=${BENCH} -benchtime=${BENCHTIME} -benchmem
