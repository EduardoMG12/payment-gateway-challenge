FROM rustlang/rust:nightly as builder
WORKDIR /usr/src/app
COPY . .
RUN cargo build --release

FROM ubuntu:24.04
WORKDIR /app
COPY --from=builder /usr/src/app/target/release/rust-processor /app/
CMD ["./rust-processor"]
