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

single-test:
	echo "Testing..."
	go test -v ${path}
