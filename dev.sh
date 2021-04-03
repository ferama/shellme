#! /bin/bash

flags="-debug=true"
# flags="$flags -ptycmd \"sh -c 'cd /tmp; exec zsh'\""

if [[ $(uname -s) == "Darwin" ]]
then
    flags="$flags -ptycmd zsh"
else
    flags="$flags -ptycmd bash"
fi

if command -v "nodemon" &> /dev/null
then
    nodemon -i ui \
        -e go \
        --exec "go run ./cmd/shellme $flags || exit 1" \
        --signal SIGTERM
else
    go run ./cmd/shellme $flags
fi