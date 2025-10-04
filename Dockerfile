# Build stage
FROM golang:1.24 AS builder
WORKDIR /goapp

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

# Final stage
FROM golang:1.24
WORKDIR /goapp
COPY --from=builder /goapp/main .
#   copy  .env
COPY --from=builder /goapp/.env .
EXPOSE 8080
CMD ["./main"]
