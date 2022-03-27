.PHONY: clean critic security lint test build run
include .env

APP_NAME = gokes
BUILD_DIR = $(PWD)/build
MIGRATIONS_FOLDER = $(PWD)/platform/migrations
DATABASE_URL = ${DB_ADDRESS_STRING}


dev:
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go
	$(BUILD_DIR)/$(APP_NAME)


clean:
	rm -rf ./build

critic:
	gocritic check -enableAll ./...

security:
	gosec ./...

lint:
	golangci-lint run ./...

# test: clean critic security lint
# 	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
# 	go tool cover -func=cover.out

build: test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go
#CGO_ENABLED=0 go build -ldfmaklags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

run: build
	$(BUILD_DIR)/$(APP_NAME)

migrate.up:	
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" up

migrate.down:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" down 1

migrate.force $(version):
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" force $(version)


migrate.create $(file): 
	migrate create -ext sql -dir $(MIGRATIONS_FOLDER) -seq $(file)




