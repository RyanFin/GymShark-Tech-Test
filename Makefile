server:
	cd backend && \
	go run main.go

test:
	cd backend && \
	go test -v ./... -coverprofile=coverage.out 
	go tool cover -html=coverage.out

run:
	cd frontend && \
	npm run dev

.PHONY: server run test
