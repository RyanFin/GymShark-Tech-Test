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

build:
	cd backend && \
	GOOS=linux GOARCH=amd64 go build -o GymShark-Tech-Test

log-fe:
	heroku logs --tail -a gymshark-tech-frontend 

log-be:
	heroku logs --tail -a gymshark-tech-backend 

.PHONY: server run swag test build log-fe log-be
