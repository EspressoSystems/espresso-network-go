lint:
    golangci-lint run ./...

test:
	success=0; \
	for i in $(seq 1 3); do \
		go test -v ./... && { success=1; break; } || { echo "Test failed. Retrying in 5 seconds..."; sleep 5; }; \
	done; \
	if [ "${success}" -eq 0 ]; then \
		echo "All test attempts failed."; \
		exit 1; \
	fi

bind-light-client:
	abigen --abi espresso-contract-artifacts/LightClient.json --pkg lightclient --out light-client/lightclient.go
	abigen --abi espresso-contract-artifacts/LightClientMock.json --pkg lightclientmock --out light-client-mock/lightclient.go

verification_dir := "./verification/rust"
target_lib := "./target/lib"

triple := if arch() == "aarch64" {
	if os() == "macos" {
		"aarch64-apple-darwin"
	} else {
		"aarch64-unknown-linux-gnu"
	}
} else if arch() == "x86_64" {
	if os() == "macos" {
		"x86_64-apple-darwin"
	} else {
		"x86_64-unknown-linux-gnu"
	}
} else {
	error("{{arch()}} is not supported")
}

build-verification:
	mkdir -p {{target_lib}}
	cargo build --release --manifest-path {{verification_dir}}/Cargo.toml
	install {{verification_dir}}/target/release/libespresso_crypto_helper.dylib {{target_lib}}/libespresso_crypto_helper-{{triple}}.dylib
	go build ./verification
