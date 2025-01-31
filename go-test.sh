#!/usr/bin/env bash

if [[ -n "$*" ]]; then
    for dir in "$@"; do
        (cd "$dir" ; go test)
    done
else
    go test
fi
