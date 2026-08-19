# Step 1: Build stage menggunakan official Go compiler image
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build statically-linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server main.go

# Step 2: Runtime image minimal (Alpine Linux, image hanya ~15MB)
FROM alpine:latest

WORKDIR /app

# Install ca-certificates & tzdata untuk TLS HTTPS & TimeZone
RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /app/server .

# Default exposed port
EXPOSE 8080

CMD ["./server"]
