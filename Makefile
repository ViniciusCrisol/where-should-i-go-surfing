.PHONY: test
test:
	go test `go list ./pkg/... | grep -v mocked/*` -p 1

.PHONY: coverage
coverage:
	go test `go list ./pkg/... | grep -v mocked/*` -p 1 -coverprofile=./coverage.out && go tool cover -html=coverage.out
