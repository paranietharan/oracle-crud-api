# Build stage
FROM golang:1.25.1 AS builder
WORKDIR /goapp

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main ./cmd/server

# Final stage
FROM golang:1.25.1
WORKDIR /goapp
COPY --from=builder /goapp/main .
#   copy  .env
COPY --from=builder /goapp/.env .
EXPOSE 8080
CMD ["./main"]
