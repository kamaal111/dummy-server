#!/bin/sh

~/go/bin/reflex -r '\.go' -s -- sh -c "PORT=8081 go run src/*.go"