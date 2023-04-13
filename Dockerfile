# stage of building
FROM golang:1.20-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
# stage of running
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY app.json .
EXPOSE 8080
CMD [ "/app/main" ]

