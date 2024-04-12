package cryptography

import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/sha256"
    "encoding/hex"
    "errors"
    "math/big"
    "time"
)

// ValidateSignedMessage validates a signed message
func ValidateSignedMessage(signedMessage map[string]interface{}, originalMessage string, nonce string, ttl int64) (bool, error) {
    // Extract components from the signed message
    ethereumAddress := signedMessage["ethereum_address"].(string)
    signature := signedMessage["signature"].(string)

    // Recover the public key from the Ethereum address
    publicKeyBytes, err := hex.DecodeString(ethereumAddress)
    if err != nil {
        return false, err
    }
    publicKey := ecdsa.PublicKey{Curve: elliptic.P256(), X: big.NewInt(0), Y: big.NewInt(0)}
    publicKey.X, publicKey.Y = publicKey.Curve.ScalarBaseMult(publicKeyBytes)

    // Verify the signature
    signatureBytes, err := hex.DecodeString(signature)
    if err != nil {
        return false, err
    }
    hashedMessage := sha256.Sum256([]byte(originalMessage))
    r := new(big.Int).SetBytes(signatureBytes[:32])
    s := new(big.Int).SetBytes(signatureBytes[32:])
    if !ecdsa.Verify(&publicKey, hashedMessage[:], r, s) {
        return false, nil
    }

    // Optionally check nonce and expiration
    if nonce != "" && signedMessage["nonce"].(string) != nonce {
        return false, errors.New("Invalid nonce")
    }
    if ttl != 0 && time.Now().Unix() > signedMessage["timestamp"].(int64)+ttl {
        return false, errors.New("Token expired")
    }

    return true, nil
}