
run: run-product

run-product:
	cd cmd/api && go mod tidy && go mod download && \
	CGO_ENABLED=0 go run main.go
.PHONY: run-product


clean:
	go clean