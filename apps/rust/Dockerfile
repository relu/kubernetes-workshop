FROM rust:1.91 AS builder

WORKDIR /app

RUN cargo init
COPY Cargo.toml Cargo.lock* ./
RUN cargo build --release && rm -rf src target/release/example-app-rust* target/release/deps/example_app_rust*

COPY src ./src

RUN cargo build --release

FROM gcr.io/distroless/cc

WORKDIR /app

COPY --from=builder /app/target/release/example-app-rust /app/example-app-rust

EXPOSE 3000

CMD ["/app/example-app-rust"]
