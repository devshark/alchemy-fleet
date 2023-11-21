.PHONY: generate
generate:
	@go generate ./ent

.PHONY: vendor
vendor:
	@go mod tidy
	@go mod vendor

.PHONY: build
build:
	CGO_ENABLED=0 go build -mod=vendor -ldflags \
		"-w -s" \
		-o bin/http \
		-tags netgo -a ./cmd

test:
	CGO_ENABLED=1 go test -race -p 10 -shuffle on ./...

# install mockery https://vektra.github.io/mockery/latest/installation/
# refer to config .mockery.yaml
mocks: vendor
	mockery
