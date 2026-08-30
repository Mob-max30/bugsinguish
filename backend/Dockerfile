# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and build binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bugsinguish-backend .

# Production stage
FROM alpine:latest

WORKDIR /app

# Install ca-certificates for HTTPS/TLS requests
RUN apk --no-cache add ca-certificates

COPY --from=builder /bugsinguish-backend .

EXPOSE 8080
CMD ["./bugsinguish-backend"]
