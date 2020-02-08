.PHONY: clean
## clean: cleans the binary and coverage file.
clean:
	rm -rf coverage.txt

.PHONY: test
## test: runs go test.
test:
	go test ./... -count=1 -race -coverprofile=coverage.txt -covermode=atomic

.PHONY: build
## build: runs go test and compiles files.
build: clean test
	go build -v ./...

.PHONY: download
## download: download Go modules dependencies.
download:
	@echo Download go.mod dependencies
	@go mod download

.PHONY: install-tools
## install-tools: installs tools imported in tools.go
install-tools: download
	@echo Installing tools from tools.go
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go get %

.PHONY: lint
## lint: running golint on files in repository.
lint: install-tools
	golint -set_exit_status ./...
