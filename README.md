# ToDo Api 
Api for todo app 

## Technologies
- Golang
- ari for hotreload 
- Docker
- Go-chi
- Gorm ORM [PostgreSQL]
- JWT Authentication
- Swagger

## Features
- [ ] Authentication
- [ ] Users Management
- [ ] Tasks Management
- [ ] Comments Management
- [ ] Tags Management
- [ ] Files Management

## Run 
```sh
$ get clone https://github.com/alaa-aqeel/todo-api.git
$ cd todo-api

# SERVER 
$ sh ./cmd.sh up  # run container 
# OR 
$ go run ./src/cmd/main.go # run server 


# RUN TEST
$ sh ./cmd.sh test  # run test in container
# OR 
$ go test ./src/test/**/*   
```

