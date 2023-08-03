#!/bin/bash

if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <commit_hash>"
    exit 1
fi

commit_hash=$1

export GOPRIVATE=github.com/Anti-Pattern-Inc/*
go get "github.com/Anti-Pattern-Inc/saasus-sdk-go@$commit_hash"
