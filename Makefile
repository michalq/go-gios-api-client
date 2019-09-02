.PHONY: all clean build view
VERSION=v0.1.1

all: clean build

build:
	$(GOPATH)/bin/swagger generate client -f ./swagger.json -A gios-api-client

clean:
	rm -rf restapi
	rm -rf models

view:
	$(GOPATH)/bin/swagger serve ./swagger.json

release:
	$(GOPATH)/bin/git-chglog --next-tag $(VERSION) -o CHANGELOG.md
	git add CHANGELOG.md
	git commit -a -m "Release $(VERSION)"
	git tag $(VERSION)