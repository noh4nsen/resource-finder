BINARY_NAME=resource-finder

build:
	mkdir build build/darwin build/linux
	GOARCH=amd64 GOOS=linux go build -o ./build/linux/${BINARY_NAME} .
	GOARCH=amd64 GOOS=darwin go build -o ./build/darwin/${BINARY_NAME} .

clean:
	go clean
	rm -rf ./build