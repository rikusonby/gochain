package handlers

import (
	"chain/blockchain"
	"encoding/hex"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func GenerateAddress(c echo.Context) error {
	chainType := c.QueryParam("type")
	var generator blockchain.AddressGenerator

	switch chainType {
	case "ethereum":
		generator = blockchain.EthereumOps{}
	default:
		err := c.String(http.StatusBadRequest, "Invalid chain type")
		return err
	}

	privateKey, _, address, err := generator.GenerateKeyPair()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to generate key pair")
	}

	// 将私钥转换为16进制字符串
	privateKeyBytes := privateKey.Bytes()
	privateKeyHex := hex.EncodeToString(privateKeyBytes)

	// 构建JSON响应
	response := map[string]string{
		"private_key": privateKeyHex,
		"address":     address,
	}

	return c.JSON(http.StatusOK, response)
}
