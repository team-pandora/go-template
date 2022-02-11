BINARY_NAME=go_template

build-app:  tidy
            CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-extldflags "-static"' -o $(BINARY_NAME) -v

clean:      go clean

run:        tidy
            go run ./main.go

tidy:       go mod tidy
