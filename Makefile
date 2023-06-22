.PHONY: build clean deploy

BINPATH := bin

build:
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ${BINPATH}/v1/addition v1/addition/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ${BINPATH}/v1/substraction v1/substraction/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ${BINPATH}/v1/multiplication v1/multiplication/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ${BINPATH}/v1/division v1/division/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ${BINPATH}/v1/squareRoot v1/squareRoot/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ${BINPATH}/v1/randomString v1/randomString/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ${BINPATH}/v1/recordList v1/record/main.go
	chmod +x bin/v1/addition
	chmod +x bin/v1/substraction
	chmod +x bin/v1/multiplication
	chmod +x bin/v1/division
	chmod +x bin/v1/squareRoot
	chmod +x bin/v1/randomString
	chmod +x bin/v1/recordList
	zip -j ${BINPATH}/v1/addition.zip ${BINPATH}/v1/addition
	zip -j ${BINPATH}/v1/substraction.zip ${BINPATH}/v1/substraction
	zip -j ${BINPATH}/v1/multiplication.zip ${BINPATH}/v1/multiplication
	zip -j ${BINPATH}/v1/division.zip ${BINPATH}/v1/division
	zip -j ${BINPATH}/v1/squareRoot.zip ${BINPATH}/v1/squareRoot
	zip -j ${BINPATH}/v1/randomString.zip ${BINPATH}/v1/randomString
	zip -j ${BINPATH}/v1/recordList.zip ${BINPATH}/v1/recordList

clean:
	rm -rf ./bin ./vendor Gopkg.lock

start-dev: clean build
	sam local start-api 

deploy: clean build
	sls deploy --verbose
