
start-network:
	$(MAKE) -C envs/ethereum-network start

deploy-contract:
	$(MAKE) -C envs/ethereum-network deploy

start:
	go run eth ./cmd/cli start

eth-run:
	go run ./cmd/cli eth run

eth-transfer:
	go run ./cmd/cli \
		eth \
		transfer \
		--to "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" \
		--from "0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"

generate:
	$(MAKE) -C envs/ethereum-network compile-go

eth-nft:
	go run ./cmd/cli \
    	eth \
    	nft --contract-path $(PWD)/envs/ethereum-network/artifacts/contracts/DoorControl.sol/DoorControl.json

run:
	go run ./cmd/cli run
