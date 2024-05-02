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

test-coverage:
	echo "Testing..."
	go test -cover ./...

single-test:
	echo "Testing..."
	go test -v ${path}

bench:
	echo "Benchmarking..."
	go test -bench=${BENCH} -benchtime=${BENCHTIME} -benchmem
