# Ethereum Wallet Generator

A simple Go application that generates Ethereum Externally Owned Account (EOA) wallets.

## Features

- Generates a new Ethereum private key
- Derives the corresponding public key
- Calculates the Ethereum address
- Saves the wallet information to a file

## Requirements

- Go 1.16 or higher
- go-ethereum package

## Installation

Clone the repository:

```bash
git clone https://github.com/YOUR_USERNAME/eth-wallet-generator.git
cd eth-wallet-generator
```
Install dependencies:
```bash
go mod download
```
## Usage
Build and run the application:
```bash
go build -o wallet-generator
./wallet-generator
```

This will:

Generate a new Ethereum wallet
Display the private key, public key, and address in the terminal
Save the wallet information to a file in the wallets directory

## Security Notes
- NEVER share your private key with anyone
- This application generates unencrypted private keys
- For production use, consider encrypting private keys or using hardware wallets
- Be careful when storing wallet information on your computer
