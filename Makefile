.PHONY: build
build:
	go build main.go

run: build
	./main

clean:
	rm -rf main