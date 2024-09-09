.PHONY: format build run

format:
	go fmt ./...

build:
	docker compose build

run: build
	docker compose up
