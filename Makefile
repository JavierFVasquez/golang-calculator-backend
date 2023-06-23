.PHONY: build clean deploy

BINPATH := bin

# List of modules and their paths
MODULES := \
	v1/addition:v1/addition/handlers/addition/main.go \
	v1/substraction:v1/substraction/handlers/substraction/main.go \
	v1/multiplication:v1/multiplication/handlers/multiplication/main.go \
	v1/division:v1/division/handlers/division/main.go \
	v1/squareRoot:v1/squareRoot/handlers/squareRoot/main.go \
	v1/randomString:v1/randomString/handlers/randomString/main.go \
	v1/recordList:v1/record/handlers/recordList/main.go \
	v1/deleteRecord:v1/record/handlers/deleteRecord/main.go \
	v1/getBalance:v1/balance/handlers/getBalance/main.go


# Rule to build a specific module
define build_module
.PHONY: build_$(word 1,$(subst :, ,$(1)))
build_$(word 1,$(subst :, ,$(1))):
	@echo "Building $(word 1,$(subst :, ,$(1)))"
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ${BINPATH}/$$(basename $(word 1,$(subst :, ,$(1)))) $(word 2,$(subst :, ,$(1)))
	chmod +x ${BINPATH}/$$(basename $(word 1,$(subst :, ,$(1))))
	zip -j ${BINPATH}/$$(basename $(word 1,$(subst :, ,$(1)))).zip ${BINPATH}/$$(basename $(word 1,$(subst :, ,$(1))))

build: build_$(word 1,$(subst :, ,$(1)))
endef

build:
$(foreach module,$(MODULES),$(eval $(call build_module,$(subst :, ,$(module)))))


clean:
	rm -rf ./bin ./vendor Gopkg.lock

start-dev: clean build
	bash env.sh
	sam local start-api

deploy: clean build
	sls deploy --verbose
