build:
	$GOOS=linux GOARCH=amd64 go build -o app_linux_amd64
	$GOOS=darwin GOARCH=arm64 go build -o app_apple_silicon
PHONY: build