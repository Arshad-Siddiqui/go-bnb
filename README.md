# Go-BnB

## Description
This is a simple project to learn Go and Gin Gonic. It is a simple API that allows you to create a user, login, and create a listing.

## How to run

### Install dependencies

```bash
go mod tidy
```
TODO: Double check the above command is accurate.

## Links to learn from
- [Link to learning file routing](https://stackoverflow.com/questions/42967235/golang-gin-gonic-split-routes-into-multiple-files)
- [Instructions used to build Dockerfile](https://hub.docker.com/_/golang)

## How to run with Docker
```bash

docker build -t go-bnb .

docker run -p 8080:8080 go-bnb

```
- The above commands will build the Docker image and run it on port 8080.