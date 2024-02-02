// Package blockchain 区块链接口定义
package blockchain

// AddressGenerator 生成地址
type AddressGenerator interface {
	GenerateAddress() string
}

// TokenBalanceFetch 获取代币余额
type TokenBalanceFetch interface {
	TokenBalance(address string) float64
}
