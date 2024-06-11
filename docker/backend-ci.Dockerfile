FROM ubuntu:22.04
WORKDIR /app
COPY backend/myapp /app
COPY backend/migrate /app/migrate
RUN chmod +x myapp

ENTRYPOINT ["./myapp"]
