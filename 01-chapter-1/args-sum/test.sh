#!/usr/bin/env bash

set -x

go build -o args-sum main.go

chmod +x args-sum

echo "test sum"
./args-sum 125.25 125.50 75.25

echo "sum with string"
./args-sum 10 "endl" 5

rm args-sum