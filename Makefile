BINARY_NAME := terragrunt

.PHONY: all $(TARGETS)

build:
	CGO_ENABLED=0
	GOARCH=amd64 GOOS=linux go build ./src -o dist/amd64/$(BINARY_NAME) -a -installsuffix cgo -ldflags="-w -s" -tags netgo -v
	GOARCH=arm64 GOOS=linux go build ./src -o dist/arm64/$(BINARY_NAME) -a -installsuffix cgo -ldflags="-w -s" -tags netgo -v
	GOARCH=arm64 GOOS=darwin go build ./src -o dist/darwin/$(BINARY_NAME) -a -installsuffix cgo -ldflags="-w -s" -tags netgo -v
