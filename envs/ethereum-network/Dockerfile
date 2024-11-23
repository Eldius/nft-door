FROM node:22-alpine

RUN mkdir /app

WORKDIR /usr/src/app

COPY . /usr/src/app
COPY docker/entrypoint.sh /usr/local/bin

RUN yarn add --dev \
    hardhat \
    "@nomicfoundation/hardhat-chai-matchers@^2.0.0" \
    "@nomicfoundation/hardhat-ethers@^3.0.0" \
    "@nomicfoundation/hardhat-ignition-ethers@^0.15.0" \
    "@nomicfoundation/hardhat-network-helpers@^1.0.0" \
    "@nomicfoundation/hardhat-verify@^2.0.0" \
    "@typechain/ethers-v6@^0.5.0" \
    "@typechain/hardhat@^9.0.0" \
    "@types/chai@^4.2.0" \
    "@types/mocha@>=9.1.0" \
    "chai@^4.2.0" \
    "ethers@^6.4.0" \
    "hardhat-gas-reporter@^1.0.8" \
    "solidity-coverage@^0.8.1" \
    "ts-node@>=8.0.0" \
    "typechain@^8.3.0" \
    "typescript@>=4.5.0"

RUN yarn install --non-interactive --frozen-lockfile

ENV NODE_OPTIONS=--max_old_space_size=4096

#ENTRYPOINT [ "npm" "hardhat" "node" ]

ENTRYPOINT ["/bin/sh", "/usr/local/bin/entrypoint.sh"]
