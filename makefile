dev: swagger
	go run ./cmd/main.go

build: swagger

swagger:
	swag init -g ./cmd/main.go -o ./docs/
