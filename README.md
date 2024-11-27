# nft-pocs #

## Configurando o projeto ##

### Go ###

- [Instalação do Go - Aprenda Go com Testes](https://larien.gitbook.io/aprenda-go-com-testes/primeiros-passos-com-go/instalacao-do-go)
- [Download and install - go.dev](https://go.dev/doc/install)

### NodeJS ###

- [Como instalar o Node.js no Windows, Linux e macOS - Alura](https://www.alura.com.br/artigos/como-instalar-node-js-windows-linux-macos)

OBS: Eu gosto de usar o NodeJS com NVM (Node Version Manager),
que permite trabalhar com multiplas versões do runtime, mas não
é algo necessário.


### Make ###

### Iniciando ###

Após configurar as ferramentas, instale as dependências
do projeto do smart contract.

```shell
# entrar na pasta do projeto
cd envs/ethereum-network

# instalar o gerenciador de pacotes Yarn
npm install -g yarn

# instalar as dependências do projeto
yarn install --non-interactive --frozen-lockfile
```

Após a configuração do projeto do smart contract
é necessário gerar os binds do contrato em Go.

```shell
# entrar na pasta do projeto
cd envs/ethereum-network

# compilar contrato
npx hardhat compile --config hardhat.config.local.js

# gerar bindings para o código Go
npx hardhat gobind --config hardhat.config.local.js
```

Daí em diante podemos rodar a aplicação em si.

## Executando o projeto ##

Tentei criar targets em um Makefile para facilitar o entendimento
das etapas.



Antes de executar o projeto localmente é necessário compilar o
contrato e iniciar a rede Ethereum.

```shell
make generate
```


## Links de referência ##

- [OpenZeppelin Contract Generator](https://wizard.openzeppelin.com/#erc721)
- [Interacting with Go bindings](https://geth.ethereum.org/docs/developers/dapp-developer/native-bindings)
- [Setup and Build Your First Web 3.0 Application With React, Hardhat, Solidity, and Metamask - dev.to](https://dev.to/suhailkakar/setup-and-build-your-first-web-30-application-with-react-hardhat-solidity-and-metamask-34jf)
