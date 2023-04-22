BINARY_NAME := terragrunt

build:
	CGO_ENABLED=0
	GOARCH=amd64 GOOS=linux go build -o ../dist/amd64/$(BINARY_NAME) -a -installsuffix cgo -ldflags="-w -s" -tags netgo -v
	GOARCH=arm64 GOOS=linux go build -o ../dist/arm64/$(BINARY_NAME) -a -installsuffix cgo -ldflags="-w -s" -tags netgo -v
	GOARCH=arm64 GOOS=darwin go build -o ../dist/darwin/$(BINARY_NAME) -a -installsuffix cgo -ldflags="-w -s" -tags netgo -v