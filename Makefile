# Generate tarball with new build of param-api-client
#
VERSION=$$(cat main.go | grep -i "cliVersion =" | awk {'print$$3'} | tr -d '"')
BINARY_NAME=psa-client

all: clean build compress report

clean:
	@rm -f /tmp/$(BINARY_NAME)-*.tar.gz
	@rm -f ./bin/$(BINARY_NAME)

build:
	@echo Building $(BINARY_NAME) version $(VERSION)
	@env CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' -o ./bin/$(BINARY_NAME)-$(VERSION)
	@cp ./bin/$(BINARY_NAME)-$(VERSION) ./bin/$(BINARY_NAME)

build-osx:
	@echo Building $(BINARY_NAME) version $(VERSION)
	@env go build -a -tags netgo -ldflags '-w' -o ./bin/$(BINARY_NAME)-osx-$(VERSION)
	@cp ./bin/$(BINARY_NAME)-osx-$(VERSION) ./bin/$(BINARY_NAME)-osx

compress:
	@tar czf /tmp/$(BINARY_NAME)-$(VERSION).tar.gz ./versions

report:
	@rm -f ./bin/$(BINARY_NAME)
	@shasum -a 256 /tmp/$(BINARY_NAME)-$(VERSION).tar.gz

.PHONY: all clean build
