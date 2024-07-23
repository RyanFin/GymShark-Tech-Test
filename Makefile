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

swag:
	cd backend && \
	swag init

.PHONY: server run swag test
