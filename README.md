# Go-BnB

## Description
This is a simple project to learn Go and ~~Gin Gonic~~ its standard net/http library, as well as GORM. It is a simple API that allows you to create a user, login, and create an airBnB style listing.

The database for this project will be powered by Postgres, and interacted with by GORM. We will also be implementing JSON Web Tokens (JWT) and bcrypt for authentication purposes.

## How to run

After cloning the repo, run the following commands in the root directory of the project.
### Install dependencies
> Note: You must have Go installed on your machine.
[Install here](https://golang.org/doc/install)

```bash
# Run in the root directory of the project
go get
```

### Run the server

```bash

go run main.go

```

## How to run with live reload

- codegangsta/gin is a live reload tool for Go web applications.
- Depending on your go version it may need to be installed globally.

```bash
# Install gin
go install github.com/codegangsta/gin

# Run the server with gin

gin run main.go

# Most likely you will need to run the following command

gin -i --appPort 8080 --port 3000 --all run main.go
# all flag will watch all files in the current directory as opposed to just go files.

```
## Links to learn from
- [File routing](https://stackoverflow.com/questions/42967235/golang-gin-gonic-split-routes-into-multiple-files)
- [Build Dockerfile](https://hub.docker.com/_/golang)
- [Official Go documentation](https://tour.golang.org/welcome/1)
- [Go by Example](https://gobyexample.com/)
- [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/)
- [Go Programming Language book](https://www.gopl.io/)
- [Go Web Examples](https://github.com/hoanhan101/go-web-examples)


## How to run with Docker

Note: You must have Docker installed on your machine.
[Install here](https://docs.docker.com/get-docker/)
```bash

docker build -t go-bnb .

docker run -dp 8080:8080 go-bnb
# Will build the Docker image and run it on port 8080.
```


## How to stop Docker container
```bash

docker ps
# Get the container id

docker stop <container id>
# Kills the container gracefully

```

If ran without -d flag, you can stop the container by pressing `ctrl+c`