FROM golang:1.22.1-alpine3.18

# Set necessary environment variables needed for our image and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN apk add --update gcc musl-dev

RUN mkdir -p /api
ADD . /api
WORKDIR /api

RUN adduser -u 10001 -D api

RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@latest && go generate . && go build -ldflags='-extldflags=-static' -o api .

CMD ["./api"]
