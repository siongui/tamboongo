ifndef GOROOT
	export GOROOT=$(realpath $(CURDIR)/../go)
	export PATH := $(GOROOT)/bin:$(PATH)
endif

commandline: fmt
	cd commandline; go run tamboon.go -rot="../fng.1000.csv.rot128"

cmdconcurrent: fmt
	go run commandline/tamboon.go -rot=fng.1000.csv.rot128 -concurrent=true -verbose=true

test: fmt
	go test -v -race

fmt:
	go fmt *.go
	go fmt commandline/*.go

modinit:
	go mod init github.com/siongui/tamboongo

modtidy:
	go mod tidy
