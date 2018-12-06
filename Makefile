TARGET=go-usb-relay

BUILD_DIR=release

ALL: release_darwin_amd64 release_windows_386
CC=i686-w64-mingw32-gcc

release_windows_386:
	@GOOS=windows GOARCH=386 CGO_ENABLED=1 CC=$(CC) go build -work -o $(BUILD_DIR)/windows/386/$(TARGET).exe

release_darwin_amd64:
	@-GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -work -o $(BUILD_DIR)/darwin/amd64/$(TARGET)

test:
	@go clean -cache
	@GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go test ./...

.PHONY: fmt
fmt:
	@go fmt ./...

.PHONY: clean
clean:
	@rm -rf $(BUILD_DIR)
