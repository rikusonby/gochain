// Package blockchain 以太坊
package blockchain

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

type EthereumOps struct{}

type EthPrivateKey struct {
	*ecdsa.PrivateKey
}

func (p EthPrivateKey) Bytes() []byte {
	return crypto.FromECDSA(p.PrivateKey)
}

type EthPublicKey struct {
	*ecdsa.PublicKey
}

func (p EthPublicKey) Bytes() []byte {
	return crypto.FromECDSAPub(p.PublicKey)
}

func (e EthereumOps) GenerateKeyPair() (PrivateKey, PublicKey, string, error) {
	privateKey, err := crypto.GenerateKey()

	if err != nil {
		return nil, nil, "", err
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()

	return EthPrivateKey{privateKey}, EthPublicKey{&privateKey.PublicKey}, address, nil
}

func (EthereumOps) Sign(data []byte, privKey PrivateKey) ([]byte, error) {
	realPrivKey, ok := privKey.(EthPrivateKey)
	if !ok {
		return nil, fmt.Errorf("invalid key type")
	}
	return crypto.Sign(crypto.Keccak256Hash(data).Bytes(), realPrivKey.PrivateKey)
}

func (EthereumOps) Verify(data []byte, sig []byte, pubKey PublicKey) bool {
	realPubKey, ok := pubKey.(EthPublicKey)
	if !ok {
		return false
	}
	hash := crypto.Keccak256Hash(data)
	sigPublicKey, err := crypto.SigToPub(hash.Bytes(), sig)
	if err != nil {
		return false
	}
	return realPubKey.PublicKey.Equal(sigPublicKey)
}
