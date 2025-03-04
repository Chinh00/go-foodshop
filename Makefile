
run: run-product

run-product:
	cd cmd/product && go mod tidy && go mod download && \
	air
.PHONY: run-product


clean:
	go clean