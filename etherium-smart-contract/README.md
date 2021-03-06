# Ethereum Smart Contract and Go

Simple Smart Contract implementation with Go following this [tutorial](https://towardsdev.com/creating-a-simple-ethereum-smart-contract-in-golang-138b9439f64e) with some tweaks.

## Setup

### Install Solidity
```console
brew update
brew upgrade
brew tap ethereum/ethereum
brew install solidity
```

### Install Geth
```console
brew install ethereum
```

### Install Ganache
https://trufflesuite.com/ganache/

--- 

## Create the Smart Contract

Create the Smart Contract: `./contracts/MySmartContract.sol`

```console
solc --optimize --abi ./contracts/MySmartContract.sol -o build

solc --optimize --bin ./contracts/MySmartContract.sol -o build

```

```console
go mod init myapp

go get -u github.com/ethereum/go-ethereum

mkdir api

abigen --abi=./build/MySmartContract.abi --bin=./build/MySmartContract.bin --pkg=api --out=./api/MySmartContract.go
```

---

## Deploy the Smart Contract

Start Ganache and take note of the **Host**, **Port**, and a **Private key**.  

Update the `/deploy/deploy.go` file with the private key and run:
```console
cd deploy
go mod tidy
go run .
```

This will produce the address where the smart contract belongs.  

---

## Communicate with the Blockchain

Update the `main.go`` file with the generated address and run it.
```console
go mod tidy
go run .
```

Try the following links in the browser:  
http://localhost:1323/hello  
http://localhost:1323/greet/bob

Or run it in the Terminal:
```console
curl http://localhost:1323/hello
curl http://localhost:1323/greet/bob
```
