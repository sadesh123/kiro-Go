.PHONY: build-all build-mac build-windows test clean

build-all: build-mac build-windows

build-mac:
	GOOS=darwin GOARCH=arm64 go build -o bin/kiro-specs-darwin-arm64 ./cmd/kiro-specs
	GOOS=darwin GOARCH=amd64 go build -o bin/kiro-specs-darwin-amd64 ./cmd/kiro-specs
	GOOS=darwin GOARCH=arm64 go build -o bin/kiro-go-darwin-arm64 ./cmd/kiro-go
	GOOS=darwin GOARCH=amd64 go build -o bin/kiro-go-darwin-amd64 ./cmd/kiro-go

build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/kiro-specs.exe ./cmd/kiro-specs
	GOOS=windows GOARCH=amd64 go build -o bin/kiro-go.exe ./cmd/kiro-go

test:
	go test ./...

clean:
	rm -rf bin/
