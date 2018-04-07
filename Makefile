all: plugins/php_plugin.so dist/server dist/plugin_validator

vendor:
	dep ensure -v

dist/server: dist vendor
	go build -o dist/server ./cmd/server

dist/plugin_validator: dist vendor
	go build -o dist/plugin_validator ./cmd/plugin/main.go

dist:
	mkdir dist

plugins/php_plugin.so: plugins vendor
	go build -buildmode=plugin -o plugins/php_plugin.so ./pkg/plugins/php/main.go

plugins:
	mkdir plugins

clean:
	rm -rf dist/*
	rm -r plugins/*

.PHONY: all clean