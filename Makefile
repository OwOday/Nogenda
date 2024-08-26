# Makefile
all: setup build

setup:
	scripts/liboleg.sh

build:
	LD_LIBRARY_PATH="$LD_LIBRARY_PATH:${PWD}/libs/OlegDB" CGO_LDFLAGS="-L/${PWD}/libs/OlegDB" go build ./src/

test: 
	LD_LIBRARY_PATH="$LD_LIBRARY_PATH:${PWD}/libs/OlegDB" CGO_LDFLAGS="-L/${PWD}/libs/OlegDB" go test -v ./src/