.PHONY: all clean build view

all: clean build

build:
	swagger generate client -f ./swagger.json -A gios-api-client

clean:
	rm -rf restapi
	rm -rf models

view:
	swagger serve ./swagger.json