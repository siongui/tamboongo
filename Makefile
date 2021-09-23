ifndef GITHUB_ACTIONS
	export GOROOT=$(realpath $(CURDIR)/../go)
	export PATH := $(GOROOT)/bin:$(PATH)
endif

test: fmt
	go test -v

fmt:
	go fmt *.go

modinit:
	go mod init github.com/siongui/tamboongo

modtidy:
	go mod tidy
