FROM rust:slim AS builder

RUN apt-get update && apt-get install -y protobuf-compiler && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY proto ./proto
COPY crypto-service ./crypto-service

WORKDIR /app/crypto-service
RUN cargo build --release

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=builder /app/crypto-service/target/release/crypto-service ./crypto-service

EXPOSE 3030

CMD ["./crypto-service"]