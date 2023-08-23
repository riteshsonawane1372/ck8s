SOURCE := $(shell find . -name '*.go')
BINARY := ck8s

build: ck8s

clean:
	@rm -rf $(BINARY)
$(BINARY) : $(SOURCE)
	CGO_ENABLED=0 go build -o $(BINARY) -ldflags="-s -w" main.go