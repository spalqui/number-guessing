.PHONY:
all: build test

build:
	go build -o ./number-guessing .

test:
	go test ./...