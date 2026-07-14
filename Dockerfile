# =============================================
# Stage 1: Build
# =============================================
FROM golang:1.23-bookworm AS builder

WORKDIR /app

# Install GCC untuk CGO (dibutuhkan go-sqlite3)
RUN apt-get update && apt-get install -y gcc musl-tools

# Copy go mod dan sum terlebih dahulu untuk cache layer
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh source code
COPY . .

# Build dengan CGO enabled (wajib untuk go-sqlite3)
RUN CGO_ENABLED=1 GOOS=linux go build -o server .

# =============================================
# Stage 2: Runtime
# =============================================
FROM debian:bookworm-slim

WORKDIR /app

# Install CA certificates dan libsqlite3
RUN apt-get update && apt-get install -y ca-certificates libsqlite3-0 && rm -rf /var/lib/apt/lists/*

# Copy binary hasil build
COPY --from=builder /app/server .

# Port yang akan dibuka
EXPOSE 8111

# Jalankan server
CMD ["./server"]
