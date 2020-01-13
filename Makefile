.init:
	git config core.hooksPath .githooks

build-go-arm: .init
	GOOS=linux GOARCH=arm GOARM=6 go build -ldflags="-s -w" -o ledctrl-example example/main.go
