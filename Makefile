default: build

# To download all dependencies
.PHONY: download_deps
download_deps:
	go mod download

# To removed unused dependencies
.PHONY: tidy
tidy:
	go mod tidy

# To vendor dependencies
.PHONY: vendor
vendor:
	go mod vendor

# To compile the go binary
.PHONY: compile
compile:
	go build -o bin/go-boilerplate -v

# To build the project, remove any unused dependencies, and compile the project
.PHONY: build
build: tidy compile

# To run the application for development
.PHONY: dev
dev:
	godotenv -f ./.dev.env go run ./main.go

# To run all the tests
.PHONY: test
test:
	godotenv -f ./.test.env  go test ./... -race -v

# To run all the tests with coverage
.PHONY: test_coverage
test_coverage:
	godotenv -f ./.test.env  go test ./... -race -v -cover
