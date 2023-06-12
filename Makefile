PROTO := $(shell find api -name "*.proto")
GO := $(shell find . -name "*.go")

run/cmd: $(GO) bin/cmd
	./bin/cmd

bin/cmd: $(GO) pb
	go build -o bin/cmd

install: $(GO) pb
	go install

.PHONY: pb clean

pb_fix:
	protodeps fix

pb: $(PROTO)
	protodeps target build golang

clean:
	rm -rf pb bin .protodeps

run_with_env:
	docker-compose up --build