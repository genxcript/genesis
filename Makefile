test:
	@go test -failfast -race -coverpkg=./... -covermode=atomic -coverprofile=coverage.txt ./... -timeout 1m