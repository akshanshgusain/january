## test: runs all tests
test:
	@go test -v ./...

## cover: opens coverage in browser
cover:
	@go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

## coverage: displays test coverage
coverage:
	@go test -cover ./...

build_cli:
	@go build -o ../januaryApp/january-cli ./cmd/cli

build_cli_desktop:
	@go build -o ~/Desktop/january-cli ./cmd/cli

count:
	 find . -type f -name '*.go' | xargs cat | wc -l

count_all:
	 find . -type f \( -name '*.go' -o -name '*.yaml' -o -name '*.jet' -o -name '*.txt' -o -name '*.sql' \) | xargs cat | wc -l

build:
	@go build -o ./dist/january ./cmd/cli
	cp ./dist/january ~/Downloads/

build_in_srv:
	@go build -o ./dist/january ./cmd/cli
	cp ./dist/january ~/Downloads/my_crud_srv/