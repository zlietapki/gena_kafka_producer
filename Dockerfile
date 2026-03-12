# Builder stage
FROM golang:1.26.1-alpine3.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./cmd/server

# Runtime stage
FROM alpine:3.23

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/server .

CMD ["./server"]
