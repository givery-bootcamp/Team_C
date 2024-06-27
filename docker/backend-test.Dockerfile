FROM golang:1.22-alpine as builder

WORKDIR /app
COPY ./backend /app/

RUN go build -cover .

#ENTRYPOINT ["GOCOVERDIR=coverdir", "./myapp"]
ENTRYPOINT ["./myapp"]
