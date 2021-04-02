#! /bin/bash

cd ui && npm install && npm run build && cd ..

build() {
    EXT=""
    [[ $GOOS = "windows" ]] && EXT=".exe"
    echo "Building ${GOOS} ${GOARCH}"
    go build -o ./bin/shellme-${GOOS}-${GOARCH}${EXT} ./cmd/shellme
}

GOOS=linux GOARCH=arm build
GOOS=linux GOARCH=arm64 build
GOOS=linux GOARCH=amd64 build

GOOS=darwin GOARCH=amd64 build
GOOS=darwin GOARCH=arm64 build