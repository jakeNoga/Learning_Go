.DEFAULT_GOAL := build

.PHONY:fmt vet build
fmt:
	go fmt ./...
vet: fmt
	go vet ./...
build: vet
	go build

.PHONY:fmt vet clean
clean:
	rm *.exe