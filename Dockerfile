# stage of building
FROM golang:1.20-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# &migrate unslash if use start.sh entrypoint for docker
# RUN apk add curl
# RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

# stage of running
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
# *migrate
# COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY .env .
COPY app.json .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./db/migration

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]

