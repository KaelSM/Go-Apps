BINARY_NAME=MarkDown.app
APP_NAME=MarkDown
VERSION=1.0.0

# build: build binary and package app
build:
	del /Q ${BINARY_NAME}
	del /Q fyne-md.exe
	fyne package -appVersion ${VERSION} -name ${APP_NAME} -release

## run: builds and runs the application
run:
	go run .

## clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@go clean
	@if exist ${BINARY_NAME} del /Q ${BINARY_NAME}
	@echo "Cleaned!"

## test: runs all tests
test:
	go test -v ./...
