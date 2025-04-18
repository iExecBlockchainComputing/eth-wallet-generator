package main

import (
	"fmt"
	"log"

	"eth-wallet-generator/internal/wallet"
)

func main() {
	// Generate a new wallet
	newWallet, err := wallet.GenerateWallet()
	if err != nil {
		log.Fatalf("Error generating wallet: %v", err)
	}

	// Print the wallet details
	fmt.Printf("Private Key: %s\n", newWallet.PrivateKey)
	fmt.Printf("Public Key: %s\n", newWallet.PublicKey)
	fmt.Printf("Ethereum Address: %s\n", newWallet.Address)

	// Save to file
	filename, err := newWallet.SaveToFile("wallets")
	if err != nil {
		log.Fatalf("Error saving wallet: %v", err)
	}
	fmt.Printf("Wallet information saved to %s\n", filename)
}
