FROM golang:1.22-alpine
WORKDIR /app
COPY backend /app

RUN go build -cover .

RUN chmod +x myapp

ENV GOCOVERDIR="coverdir/e2e"

RUN chmod +x myapp.sh

ENTRYPOINT ["./myapp.sh"]
