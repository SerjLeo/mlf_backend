# My Local Financier Backend

## Technology stack

- [go](https://go.dev/) 1.16
- [gin](https://github.com/gin-gonic/gin) 1.7.4
- [docker](https://www.docker.com/) 20.10.13
- [docker-compose](https://docs.docker.com/compose/) 1.29.2

## Database migrations
Postgres 14+ is used

Migrations located in _migrations_ folder

To run migrations you can use [go-migrate](https://github.com/golang-migrate/migrate). After installation use command bellow in root directory:

`migrate -path ./migrations -database postgresql://{username}:{password}@{host}:{port}/{dbname}?sslmode=disable up`

## Run project with Docker Compose

To build app containers [docker-compose](https://docs.docker.com/compose/) is used.

Initial build command:

`docker-compose up -d`

To update server on running container:

`docker-compose up --no-deps --force-recreate -d --build server`