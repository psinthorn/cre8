BINARY_NAME=cre8App

build:
	@go mod vendor
	@echo "Building Cre8..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Cre8 built!"

run: build
	@echo "Starting Cre8..."
	@./tmp/${BINARY_NAME} &
	@echo "Cre8 started!"

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
	@echo "Stopping Cre8..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped Cre8!"

restart: stop start