build:
	echo "Building..."
	go build -o $(path)/main.exe $(path)/main.go

run:
	echo "Running..."
	go run $(path)
