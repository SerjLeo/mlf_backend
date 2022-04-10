# My Local Financier Backend

## Database migrations
Migrations located in _migrations_ folder

To run migrations you can use [go-migrate](https://github.com/golang-migrate/migrate). After installation use command bellow in root directory:

`migrate -path ./migrations -database postgresql://{username}:{password}@{host}:{port}/{dbname}?sslmode=dasable up`

## Run project with Docker Compose

To build app containers [docker-compose](https://docs.docker.com/compose/) is used.

Initial build command:

`docker-compose up -d`

To update server on running container:

`docker-compose up --no-deps --force-recreate -d --build server`