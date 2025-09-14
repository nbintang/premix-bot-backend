.PHONY: run build clean

dev:
	air --build.cmd "go build -o tmp/server.exe cmd/server/main.go" --build.bin "tmp\\server.exe"

build:
	go build -o bin/app .

clean:
	rm -rf bin tmp
