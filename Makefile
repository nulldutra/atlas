BINARY_NAME=atlas

benchmark: plow http://localhost:8000 -c 2 --requests=20

build: go build -o ./bin/${BINARY_NAME}

run: go run main.go

default: benchmark build run
