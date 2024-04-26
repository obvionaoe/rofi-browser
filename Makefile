BINARY_NAME = rofi-browser

build:
	go build -v

run: build
	./$(BINARY_NAME)

install: build
	sudo cp ./$(BINARY_NAME) /usr/bin/
