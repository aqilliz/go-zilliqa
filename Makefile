# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: gzil ios gzil-cross evm all test clean
.PHONY: gzil-linux gzil-linux-386 gzil-linux-amd64 gzil-linux-mips64 gzil-linux-mips64le
.PHONY: gzil-linux-arm gzil-linux-arm-5 gzil-linux-arm-6 gzil-linux-arm-7 gzil-linux-arm64
.PHONY: gzil-darwin gzil-darwin-386 gzil-darwin-amd64
.PHONY: gzil-windows gzil-windows-386 gzil-windows-amd64

GOBIN = $(shell pwd)/build/bin
GO ?= latest

gzil:
	build/env.sh go run build/ci.go install ./cmd/gzil
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gzil\" to launch gzil."

all:
	build/env.sh go run build/ci.go install

test: all
	build/env.sh go run build/ci.go test

lint: ## Run linters.
	build/env.sh go run build/ci.go lint

clean:
	./build/clean_go_build_cache.sh
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go get -u golang.org/x/tools/cmd/stringer
	env GOBIN= go get -u github.com/kevinburke/go-bindata/go-bindata
	env GOBIN= go get -u github.com/fjl/gencodec
	env GOBIN= go get -u github.com/golang/protobuf/protoc-gen-go
	env GOBIN= go install ./cmd/abigen
	@type "npm" 2> /dev/null || echo 'Please install node.js and npm'
	@type "solc" 2> /dev/null || echo 'Please install solc'
	@type "protoc" 2> /dev/null || echo 'Please install protoc'

# Cross Compilation Targets (xgo)

gzil-cross: gzil-linux gzil-darwin gzil-windows gzil-android gzil-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/gzil-*

gzil-linux: gzil-linux-386 gzil-linux-amd64 gzil-linux-arm gzil-linux-mips64 gzil-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/gzil-linux-*

gzil-linux-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/gzil
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/gzil-linux-* | grep 386

gzil-linux-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/gzil
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gzil-linux-* | grep amd64

gzil-linux-arm: gzil-linux-arm-5 gzil-linux-arm-6 gzil-linux-arm-7 gzil-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/gzil-linux-* | grep arm

gzil-linux-arm-5:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-5 -v ./cmd/gzil
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/gzil-linux-* | grep arm-5

gzil-linux-arm-6:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-6 -v ./cmd/gzil
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/gzil-linux-* | grep arm-6

gzil-linux-arm-7:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-7 -v ./cmd/gzil
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/gzil-linux-* | grep arm-7

gzil-linux-arm64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm64 -v ./cmd/gzil
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/gzil-linux-* | grep arm64

gzil-linux-mips:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/gzil
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/gzil-linux-* | grep mips

gzil-linux-mipsle:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/gzil
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/gzil-linux-* | grep mipsle

gzil-linux-mips64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/gzil
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/gzil-linux-* | grep mips64

gzil-linux-mips64le:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/gzil
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/gzil-linux-* | grep mips64le

gzil-darwin: gzil-darwin-386 gzil-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/gzil-darwin-*

gzil-darwin-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/gzil
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/gzil-darwin-* | grep 386

gzil-darwin-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/gzil
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gzil-darwin-* | grep amd64

gzil-windows: gzil-windows-386 gzil-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/gzil-windows-*

gzil-windows-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/386 -v ./cmd/gzil
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/gzil-windows-* | grep 386

gzil-windows-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/gzil
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/grep-windows-* | grep amd64
