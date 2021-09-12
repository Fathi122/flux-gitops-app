FROM golang:alpine

RUN mkdir /app

COPY  ./go.mod /app
COPY  ./go.sum /app
COPY  ./appserver.go /app

WORKDIR /app

RUN go build -o main .

EXPOSE 8080
CMD ["/app/main"]