# Go-BnB

## Description
This is a simple project to learn Go and ~~Gin Gonic~~ its standard net/http library. It is a simple API that allows you to create a user, login, and create an airBnB style listing.

The database for this project will be powered by MongoDB, and we will be using the official Go driver to connect to it. We will also be implementing JSON Web Tokens (JWT) and bcrypt for authentication purposes.

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

## Links to learn from
- [File routing](https://stackoverflow.com/questions/42967235/golang-gin-gonic-split-routes-into-multiple-files)
- [Build Dockerfile](https://hub.docker.com/_/golang)

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