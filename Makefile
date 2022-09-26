build:
	go build -o ./outputs/out src/*

run: build
	./outputs/out
