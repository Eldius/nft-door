
start:
	npx hardhat --config hardhat.config.local.js node

deploy:
	npx hardhat run scripts/deploy.js --network hardhat

compile:
	npx hardhat compile --config hardhat.config.local.js

compile-go: compile
	abigen --abi Storage.abi --pkg main --type Storage --out Storage.go --bin Storage.bin
	npx hardhat gobind --config hardhat.config.local.js
