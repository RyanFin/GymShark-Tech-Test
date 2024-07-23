run:
	go run main.go
.PHONY: run

test:
	go test -v ./... -coverprofile=coverage.out 
	go tool cover -html=coverage.out
