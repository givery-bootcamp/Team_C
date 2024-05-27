FROM golang:1.22.2-alpine
RUN apk add --update --no-cache ca-certificates git

WORKDIR /go/src/myapp

RUN go install github.com/cosmtrek/air@v1.51.0


ENV DOCKERIZE_VERSION v0.7.0

RUN apk update --no-cache \
    && apk add --no-cache wget openssl \
    && wget -O - https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz | tar xzf - -C /usr/local/bin \
    && apk del wget

ENTRYPOINT dockerize -timeout 10s -wait tcp://db:3306  air