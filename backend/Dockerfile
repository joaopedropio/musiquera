# Stage 1: Build
FROM --platform=$BUILDPLATFORM golang:1.24-alpine AS builder

WORKDIR /app
COPY ./backend/go.mod ./backend/go.sum ./
RUN go mod download
COPY ./backend .

RUN GOOS=linux GOARCH=arm64 go build -o /app/server .

# Stage 2: Runtime
FROM --platform=linux/arm64 alpine:latest

WORKDIR /app

COPY --from=builder /app/server ./server

COPY ./frontend/dist /static

EXPOSE 8080

ENTRYPOINT ["./server"]

