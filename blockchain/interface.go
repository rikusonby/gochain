// Package blockchain 区块链接口定义
package blockchain

type PrivateKey interface {
	Bytes() []byte
}

type PublicKey interface {
	Bytes() []byte
}

// AddressGenerator 生成地址
type AddressGenerator interface {
	GenerateKeyPair() (PrivateKey, PublicKey, string, error)
	Sign(data []byte, privKey PrivateKey) ([]byte, error)
	Verify(data []byte, sig []byte, pubKey PublicKey) bool
}

// TokenBalanceFetch 获取代币余额
type TokenBalanceFetch interface {
	TokenBalance(address string) float64
}
