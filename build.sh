#! /bin/bash

cd ui && npm install && npm run build && cd ..

go build -o ./bin/shellme ./cmd/shellme