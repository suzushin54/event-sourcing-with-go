

#build-debug:
#	type dlv || go install github.com/go-delve/delve/cmd/dlv@v1.9.0
#	mkdir -p tmp
#	go build -gcflags "all=-N -l" -o ./tmp/main ./cmd/main

.PHONY: wire
wire: ## wire
	wire ./cmd/di
