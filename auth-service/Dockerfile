FROM golang:alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /app

COPY auth-service/go.mod auth-service/go.sum ./auth-service/

WORKDIR /app/auth-service
RUN go mod download

COPY auth-service/ ./

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o auth-app ./cmd/server/main.go

FROM alpine:3.20

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/auth-service/auth-app .

EXPOSE 3131

CMD ["./auth-app"]