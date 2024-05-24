PLATFORM=$(shell uname -s | tr '[:upper:]' '[:lower:]')
VERSION := $(shell grep -Eo '(v[0-9]+[\.][0-9]+[\.][0-9]+(-[a-zA-Z0-9]*)?)' version.go)

.PHONY: build build-server docker release check

build:
	CGO_ENABLED=0 go build -o ./bin/server github.com/moov-io/wire/cmd/server

build-webui:
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js ./docs/webui/assets/wasm_exec.js
	GOOS=js GOARCH=wasm go build -o ./docs/webui/assets/wire.wasm github.com/moov-io/wire/docs/webui/

.PHONY: check
check:
ifeq ($(OS),Windows_NT)
	@echo "Skipping checks on Windows, currently unsupported."
else
	@wget -O lint-project.sh https://raw.githubusercontent.com/moov-io/infra/master/go/lint-project.sh
	@chmod +x ./lint-project.sh
	GOOS=js GOARCH=wasm GOCYCLO_LIMIT=115 COVER_THRESHOLD=55.0 time ./lint-project.sh
endif

.PHONY: client
client:
# Versions from https://github.com/OpenAPITools/openapi-generator/releases
	@chmod +x ./openapi-generator
	@rm -rf ./client
	OPENAPI_GENERATOR_VERSION=4.2.0 ./openapi-generator generate -i openapi.yaml -g go -o ./client
	rm -f client/go.mod client/go.sum ./client/.travis.yml
	go fmt ./...
	go build github.com/moov-io/wire/client
	go test ./client

.PHONY: clean
clean:
	@rm -rf ./bin/ ./tmp/ coverage.txt misspell* staticcheck lint-project.sh

dist: clean client build
ifeq ($(OS),Windows_NT)
	CGO_ENABLED=1 GOOS=windows go build -o bin/wire.exe github.com/moov-io/wire/cmd/server
else
	CGO_ENABLED=0 GOOS=$(PLATFORM) go build -o bin/wire-$(PLATFORM)-amd64 github.com/moov-io/wire/cmd/server
endif

dist-webui: build-webui
	git config user.name "moov-bot"
	git config user.email "oss@moov.io"
	git add ./docs/webui/assets/wasm_exec.js ./docs/webui/assets/wire.wasm
	git commit -m "chore: updating wasm webui" || echo "No changes to commit"
	git push origin master

docker: clean docker-hub docker-openshift

docker-hub:
	docker build --pull -t moov/wire:$(VERSION) -f Dockerfile .
	docker tag moov/wire:$(VERSION) moov/wire:latest

docker-openshift:
	docker build --pull -t quay.io/moov/wire:$(VERSION) -f Dockerfile.openshift --build-arg VERSION=$(VERSION) .
	docker tag quay.io/moov/wire:$(VERSION) quay.io/moov/wire:latest

.PHONY: clean-integration test-integration

clean-integration:
	docker-compose kill
	docker-compose rm -v -f

test-integration: clean-integration
	docker-compose up -d
	sleep 5
	curl -v http://localhost:8088/files

release: docker AUTHORS
	go vet ./...
	go test -coverprofile=cover-$(VERSION).out ./...
	git tag -f $(VERSION)

release-push:
	docker push moov/wire:$(VERSION)
	docker push moov/wire:latest

quay-push:
	docker push quay.io/moov/wire:$(VERSION)
	docker push quay.io/moov/wire:latest

.PHONY: cover-test cover-web
cover-test:
	go test -coverprofile=cover.out ./...
cover-web:
	go tool cover -html=cover.out

# From https://github.com/genuinetools/img
.PHONY: AUTHORS
AUTHORS:
	@$(file >$@,# This file lists all individuals having contributed content to the repository.)
	@$(file >>$@,# For how it is generated, see `make AUTHORS`.)
	@echo "$(shell git log --format='\n%aN <%aE>' | LC_ALL=C.UTF-8 sort -uf)" >> $@

.PHONY: tagged-release
tagged-release:
	@./tagged-release.sh $(VERSION)
