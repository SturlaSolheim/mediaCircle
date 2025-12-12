#!/bin/bash

# Dette scriptet genererer mocks og kjører alle tester

set -e  

echo "Genererer mocks"
rm -r mocks
go run github.com/vektra/mockery/v2@latest

echo "Mocks generert"

echo "Kjører tester"
go test ./...

