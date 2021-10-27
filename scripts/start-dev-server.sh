#!/bin/sh

~/go/bin/reflex -r '\.go' -s -- sh -c "PORT=8000 go run src/*.go"