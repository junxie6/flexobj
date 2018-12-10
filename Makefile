SHELL = /bin/bash

ROOT_DIR = $(CURDIR)

APP_NAME = github.com/junxie6/flexobj

go-test:
	set -o pipefail; cd $(ROOT_DIR) && go test -v -mod vendor -race $(APP_NAME)/...

go-test-set:
	@set -o pipefail; cd $(ROOT_DIR) && go test -v -mod vendor -race -run="Set" $(APP_NAME)/...

go-test-decode:
	@set -o pipefail; cd $(ROOT_DIR) && go test -v -mod vendor -race -run="Decode" $(APP_NAME)/...

go-test-clone:
	@set -o pipefail; cd $(ROOT_DIR) && go test -v -mod vendor -race -run="Clone" $(APP_NAME)/...

go-test-cover:
	set -o pipefail; cd $(ROOT_DIR) && go test -v -cover -mod vendor $(APP_NAME)/...

go-tidy:
	set -o pipefail; cd $(ROOT_DIR) && go mod tidy -v

go-clean:
	set -o pipefail; cd $(ROOT_DIR) && go clean -i -x -modcache

