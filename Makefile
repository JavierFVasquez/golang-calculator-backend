.PHONY: build clean deploy

BINPATH := bin

build:
	env GOOS=linux go build -ldflags="-s -w" -o ${BINPATH}/v1/addition v1/addition/main.go
	env GOOS=linux go build -ldflags="-s -w" -o ${BINPATH}/v1/subtraction v1/subtraction/main.go
	chmod +x bin/v1/addition
	chmod +x bin/v1/subtraction
	zip -j ${BINPATH}/v1/addition.zip ${BINPATH}/v1/addition
	zip -j ${BINPATH}/v1/subtraction.zip ${BINPATH}/v1/subtraction

clean:
	rm -rf ./bin ./vendor Gopkg.lock

start-dev: clean build
	sam local start-api 

deploy: clean build
	sls deploy --verbose
