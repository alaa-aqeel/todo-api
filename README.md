# ToDo 

## Features
- [ ] Authentication 
- [ ] Authorization
- [X] Users Management
- [ ] Tasks Management
- [ ] Comments Management
- [ ] Tags Management
- [ ] Files Management

## Technologies
- Golang
- Docker
- Nginx
- PostgresSQL 
- JWT Authentication

## Run 
```sh
$ get clone https://github.com/alaa-aqeel/todo-api.git
$ cd todo-api

# SERVER 
$ sh ./cmd.sh up  # run container 
# OR 
$ go run ./src/cmd/main.go # run server 


# RUN TEST
$ sh ./cmd.sh test ./... #OR units/... | features/... 
# OR 
$ go test ./src/test/... 
```

