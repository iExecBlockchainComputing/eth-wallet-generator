package wallet

import (
	"crypto/ecdsa"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// Wallet represents an Ethereum wallet
type Wallet struct {
	PrivateKey    string
	PublicKey     string
	Address       string
	TimeGenerated time.Time
}

// GenerateWallet creates a new Ethereum wallet
func GenerateWallet() (*Wallet, error) {
	// Generate a new private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %v", err)
	}

	// Get the private key in hex format
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := hexutil.Encode(privateKeyBytes)

	// Generate the public key from the private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed to cast public key to ECDSA")
	}

	// Get the public key in hex format
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	publicKeyHex := hexutil.Encode(publicKeyBytes)

	// Generate the Ethereum address from the public key
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	addressHex := address.Hex()

	return &Wallet{
		PrivateKey:    privateKeyHex,
		PublicKey:     publicKeyHex,
		Address:       addressHex,
		TimeGenerated: time.Now(),
	}, nil
}

// SaveToFile stores wallet information in a file
func (w *Wallet) SaveToFile(directory string) (string, error) {
	// Create a directory to store the wallet if it doesn't exist
	if err := os.MkdirAll(directory, 0755); err != nil {
		return "", fmt.Errorf("failed to create directory: %v", err)
	}

	// Format the wallet information
	timestamp := w.TimeGenerated.Format("2006-01-02_15-04-05")
	walletInfo := fmt.Sprintf("Wallet generated on: %s\nPrivate Key: %s\nPublic Key: %s\nEthereum Address: %s\n",
		timestamp, w.PrivateKey, w.PublicKey, w.Address)

	// Save the wallet information to a file
	filename := filepath.Join(directory, fmt.Sprintf("wallet_%s.txt", timestamp))
	if err := os.WriteFile(filename, []byte(walletInfo), 0644); err != nil {
		return "", fmt.Errorf("failed to save wallet information: %v", err)
	}

	return filename, nil
}
