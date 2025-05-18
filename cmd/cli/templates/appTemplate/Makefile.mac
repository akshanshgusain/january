BINARY_NAME=januaryApp

build:
	@echo "Building January..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "January built!"

run: build
	@echo "Starting January..."
	@./tmp/${BINARY_NAME} &
	@echo "January started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping January..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped January!"

restart: stop start

start_db:
	@echo "Starting DBs with docker-compose"
	docker compose -f ./database/docker-compose.yml -p january up -d

stop_db:
	@echo "Stopping DBs with docker-compose"
	docker compose -f ./database/docker-compose.yml -p january down

count:
	 find . -type f -name '*.go' | xargs cat | wc -l

count_all:
	 find . -type f \( -name '*.go' -o -name '*.yaml' -o -name '*.jet' \) | xargs cat | wc -l