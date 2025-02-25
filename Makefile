vine_path = $(HOME)/.vine
vine_name = vine_aly

.PHONY: help
help:
    @echo 'Usage:'

.PHONY: build
build:
	go build -o bin/$(vine_name) main.go

.PHONY: clean
clean:
	rm -f $(vine_path)/$(vine_name)
	rm -f bin/$(vine_name)

.PHONY: install
install: clean build
	cp bin/$(vine_name) $(vine_path)

