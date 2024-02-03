BIN_NAME=server-exe
DEBUG=true

run: 
	@echo "Начинаю запускать сервер через makefile"
	@go run ./cmd/api

dev:
	@DEBUG=$(DEBUG) go run ./cmd/api

.PHONY: build
build: 
	go build -o $(BIN_NAME) ./cmd/api

run-after-build: build
	./$(BIN_NAME)

clean:
	rm $(BIN_NAME)