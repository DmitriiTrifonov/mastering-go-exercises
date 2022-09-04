#!/usr/bin/env bash

set -x

go build -o args-avg main.go

chmod +x args-avg

echo "test avg"
./args-avg 2.25 4.50 2.25

echo "avg with string"
./args-avg 5 "endl" 5

rm args-avg