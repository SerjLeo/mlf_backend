# My Local Financier Backend

## Technology stack

- [go](https://go.dev/) 1.16
- [gin](https://github.com/gin-gonic/gin) 1.7.4
- [docker](https://www.docker.com/) 20.10.13
- [docker-compose](https://docs.docker.com/compose/) 1.29.2

## Database migrations
Postgres 14+ is used

Migrations located in _migrations_ folder

To run migrations manually you can use [go-migrate](https://github.com/golang-migrate/migrate). After installation use command bellow in root directory:

`migrate -path ./migrations -database postgresql://{username}:{password}@{host}:{port}/{dbname}?sslmode=disable up`

To run database container:

`docker run --name mlf-db -e POSTGRES_USER=user -e POSTGRES_DB=db -e POSTGRES_PASSWORD=pass -d -p 5432:5432 postgres`

## Documentation generation

For API documentation [swag](https://github.com/swaggo/swag) library is used. To generate swagger docs install swag cli and use command

`swag init -g cmd/app/main.go`

API specification will be available at http://localhost:port/swagger/index.html.

## Testing

For testing [testify](https://github.com/stretchr/testify) and [mockery](https://github.com/vektra/mockery) is used. To generate mocks use mockery command:

`mockery`

## Run project with Docker Compose

To build app containers [docker-compose](https://docs.docker.com/compose/) libraries is used.

Initial build command:

`docker-compose up -d`

To update server on running container:

`docker-compose up --no-deps --force-recreate -d --build server`