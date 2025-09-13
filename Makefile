.PHONY: run build clean

dev:
	air -c .air.toml

build:
	go build -o bin/app .

clean:
	rm -rf bin tmp
