ifndef GITHUB_ACTIONS
	export GOROOT=$(realpath $(CURDIR)/../go)
	export PATH := $(GOROOT)/bin:$(PATH)
endif

test:
	go test -v

modinit:
	go mod init github.com/siongui/tamboongo

modtidy:
	go mod tidy
