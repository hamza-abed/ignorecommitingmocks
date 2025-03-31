
mock:
	go run github.com/vektra/mockery/v2@latest --name UserService --dir internal/service --output mocks --outpkg mocks

clean:
	rm -rf mocks
	rm app

build:
	GOOS=linux GOARCH=arm64 GOAMD64=v3 CGO_ENABLED=0 go build -o app ./cmd

test:
	go test -v ./...
