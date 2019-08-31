.PHONY: all clean build view

all: clean build

build:
	swagger generate client -f ./swagger.json -A gios-api-client


 generate client -f /home/michal/www/gios-api-client/swagger.json -A gios-api-client
clean:
	rm -rf restapi
	rm -rf models

view:
	swagger serve ./swagger.json

release:
	$GOPATH/bin/git-chglog 