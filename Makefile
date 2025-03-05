run:
	@sh -c 'go run $$(find cmd/web -type f -name "*.go" ! -name "*_test.go")'