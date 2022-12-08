# gqlgen-todos
This Repository is using gqlgen as a trial

## System architecture
```mermaid
graph LR
    front --> |query/mutation|api[galgen<br>GraphQL]
    api --> mysql
    api <-->|pub/sub| redis
    api --> |subscription|front
```

## Requirement
### api

* Go 1.19.*
* gqlgen v0.17.*
* gin 1.8.* (web framework)
* gorm v2 1.24.* (orm)
* golang-migrate V4.15.*(migrate)
* wire v0.5.*(di)

### front

* typescript 4.9.*
* next 13.0.*
* urql 3.0.*

## Usage

start
```bash
$ make start
```

graphql server build([gqlgen](https://gqlgen.com/))
```bash
$ make exec api
## in container
make generate
```

di generate ([wire](https://github.com/google/wire))
```bash
$ make exec api
## in container
$ make di
```

db migrate ([golang-migrate](https://github.com/golang-migrate/migrate))
```bash
$ make exec api
## in container
### create migration template
$ migrate create -ext sql -dir db/migrations -seq create_todos_table
### migration up
$ migrate -source file://db/migrations -database "mysql://$DB_USERNAME:$DB_PASSWORD@tcp($DB_HOST:$DB_PORT)/$DB_DATABASE" up
### migration up
$ migrate -source file://db/migrations -database "mysql://$DB_USERNAME:$DB_PASSWORD@tcp($DB_HOST:$DB_PORT)/$DB_DATABASE" down
```