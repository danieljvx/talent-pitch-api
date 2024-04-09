FROM golang:alpine as builder

LABEL maintainer="Daniel Villanueva <villanueva.danielx@gmail.com>"

RUN apk update && apk add --no-cache git

# Set necessary environment variables needed for our image and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN apk add --update gcc musl-dev

RUN mkdir -p /api
ADD . /api
WORKDIR /api

RUN adduser -u 10001 -D api

RUN go mod download

RUN go get -u github.com/swaggo/swag/cmd/swag

RUN go install github.com/swaggo/swag/cmd/swag@latest

# RUN swag init -g main.go --output docs

RUN go generate . && go build -ldflags='-s -w -extldflags=-static' -o api .

CMD ["./api"]
