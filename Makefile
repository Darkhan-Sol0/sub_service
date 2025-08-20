NAME = service

EXEC = app

SOURCE = ./cmd/app/main.go

PACKAGE =\
					github.com/ilyakaznacheev/cleanenv\
					github.com/labstack/echo/v4\
					github.com/jackc/pgx/v5\
					github.com/jackc/pgx/v5/pgxpool\
					github.com/swaggo/http-swagger\
					github.com/alecthomas/template\
					github.com/swaggo/echo-swagger\
					github.com/sirupsen/logrus\

PACKAGE_GOOSE =\
					github.com/pressly/goose/v3/cmd/goose@latest\

PACKAGE_SWAGGER =\
					github.com/swaggo/swag/cmd/swag@latest

.PHONY: all build run clean init get get_install docker_clean docker

all: clean build run

build:
	go build -o $(EXEC) $(SOURCE)

run:
	./$(EXEC)

clean:
	rm -f $(EXEC)

init:
	go mod init $(NAME)

get:
	go get -u $(PACKAGE)

goose_up:
	~/go/bin/goose -dir migrations postgres "postgresql://user:password@pgdb:5432/db?sslmode=disable" up

goose_down:
	~/go/bin/goose -dir migrations postgres "postgresql://user:password@pgdb:5432/db?sslmode=disable" down

goose_install:
	go install $(PACKAGE_GOOSE)

swagger_install:
	go install $(PACKAGE_SWAGGER)

swagger_init:
	~/go/bin/swag init -g cmd/app/main.go

docker_clean:
	docker rm app migrations postgres
	docker rmi tz_project-migrations tz_project-app

docker:
	docker compose up