.PHONY: all build run test check clean

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

clean:
	cargo clean
