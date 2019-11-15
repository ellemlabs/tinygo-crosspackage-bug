TINYGO = /usr/local/tinygo/bin/tinygo
GO     = /usr/local/go/bin/go


# Fails:
build.tinygo.wasm: main-test.go
	$(TINYGO) build -target=wasm -o build.tinygo.wasm main-test.go

# Works:
build.wasm: main-test.go
	GOOS=js GOARCH=wasm $(GO) build -o build.wasm main-test.go
	$(GO) build
