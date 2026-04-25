.PHONY: all build run test check format clean

all: build

build:
	cargo build

run:
	cargo run

test:
	cargo test

check:
	cargo fmt --check
	cargo clippy -- -D warnings

format:
	cargo fmt

clean:
	cargo clean
