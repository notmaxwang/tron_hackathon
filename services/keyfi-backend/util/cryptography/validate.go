package cryptography

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/keystore"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

const (
	DEFAULT_MESSAGE = "Hello, I am logging into KeyFi!"
	TronBytePrefix  = byte(0x41)
)

func ValidateDefaultMessage(expiry int64, signatureRaw string, walletAddress string) (bool, error) {
	message := fmt.Sprintf("%s %v", DEFAULT_MESSAGE, expiry)
	return ValidateMessage(message, signatureRaw, walletAddress)
}

func ValidateMessage(messageRaw string, signatureRaw string, walletAddress string) (bool, error) {
	// Convert the signature to bytes
	sigBytes, err := hex.DecodeString(signatureRaw[2:])
	if err != nil {
		log.Printf("Error decoding signature:", err)
		return false, err
	}

	// Hash the message
	hash := keystore.TextHash([]byte(messageRaw), true)

	// V value needs to be capped
	if sigBytes[64] >= 27 {
		sigBytes[64] -= 27
	}

	// Recover the Ethereum address from the signature
	sigPublicKey, err := secp256k1.RecoverPubkey(hash, sigBytes)
	if err != nil {
		log.Printf("error while recovering key", err)
		return false, err
	}
	pubKey, err := keystore.UnmarshalPublic(sigPublicKey)
	if err != nil {
		log.Printf("error while unmarhsalling", err)
		return false, err
	}

	// Convert the recovered public key to an Ethereum address
	recoveredSignerAddress := address.PubkeyToAddress(*pubKey)

	walletAddressDecoded, err := address.Base58ToAddress(walletAddress)
	if err != nil {
		log.Printf("error while decoding wallet address", err)
		return false, err
	}

	log.Printf("\noriginal: %s\nrecovered: %s\n", walletAddressDecoded, recoveredSignerAddress)
	// Compare the recovered address with the provided signer address
	return bytes.Equal(walletAddressDecoded, recoveredSignerAddress), nil
}
