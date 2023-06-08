DOCKER_IMAGE_NAME="fiber-cookiecutter"
DOCKER_IMAGE_TAG="latest"

run:
	go run main.go
test:
	go test -v ./...

build:
	docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG) .

up:
	docker-compose up -d 

down:
	docker-compose down

swag:
	swag init

.PHONY: run test build