BINARY=zx

.PHONY: build
build: tidy clean
	GOOS=linux go build -o ./$(BINARY)

.PHONY: clean
clean:
	go clean

.PHONY: install
install: build
	sudo cp $(BINARY) /usr/local/bin

.PHONY: run
run: build
	./$(BINARY)

.PHONY: test
test:
	go test ./... -race -cover -count=1

.PHONY: tidy 
tidy:
	go mod tidy
