BINARY=logmonitor

lint:
	golangci-lint run

run:
	go run cmd/logmonitor/main.go

test:
	go test ./internal/monitor -v

build:
	go build -o $(BINARY) ./cmd/logmonitor

clean:
	rm -rf $(BINARY) output/


help:
	@echo ""
	@echo "\tmake lint    - lint the project"
	@echo "\tmake run     - run the project"
	@echo "\tmake test    - run tests"
	@echo "\tmake build   - build the project"
	@echo "\tmake clean   - clean binaries and reports"
	@echo ""
