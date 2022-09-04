#!/usr/bin/env bash

set -x

go build -o read-int main.go

chmod +x read-int

echo "test args"
./read-int 1 2 3 4 5 6 7 8 9 0 0 0 END

echo "with string"
./read-int 10 "endl" 5 END

echo "nums.txt"
./read-int < nums.txt

rm read-int