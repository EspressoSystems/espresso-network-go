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



os_name := `uname -s`

build-verification:
	@if [ "{{os_name}}" != "Darwin" ]; then \
		export LD_LIBRARY_PATH="$$(pwd)/target/lib:$$LD_LIBRARY_PATH"; \
	fi
	mkdir -p {{target_lib}}
	cargo build --release --manifest-path {{verification_dir}}/Cargo.toml
	@if [ "{{os_name}}" == "Darwin" ]; then \
		install {{verification_dir}}/target/release/libespresso_crypto_helper.dylib {{target_lib}}/libespresso_crypto_helper.dylib; \
	else \
		install {{verification_dir}}/target/release/libespresso_crypto_helper.so {{target_lib}}/libespresso_crypto_helper.so; \
	fi
	go build ./verification