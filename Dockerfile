# Build stage
FROM golang:1.25-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

# Set working directory
WORKDIR /build

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build arguments
ARG BUILD_VERSION=unknown
ARG BUILD_DATE=unknown

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s -X 'main.Version=${BUILD_VERSION}' -X 'main.BuildDate=${BUILD_DATE}'" \
    -o backoffice-service \
    ./cmd/main.go

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates tzdata

# Create non-root user
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser

# Create necessary directories
RUN mkdir -p /app/storage/logs && \
    chown -R appuser:appuser /app

# Set working directory
WORKDIR /app

# Copy binary from builder
COPY --from=builder /build/backoffice-service /app/backoffice-service

# Copy configuration files if needed
# COPY --from=builder /build/config ./config

# Change ownership
RUN chown -R appuser:appuser /app && \
    chmod +x /app/backoffice-service

# Switch to non-root user
USER appuser

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Run the application
CMD ["./backoffice-service"]

