FROM golang:1.21.3-alpine3.18

WORKDIR /twitter-clone/server

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main ./cmd

EXPOSE 8080

CMD ["./main"]