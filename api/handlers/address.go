package handlers

import (
	"chain/blockchain"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GenerateAddress(c echo.Context) error {
	chainType := c.QueryParam("type")
	var generator blockchain.AddressGenerator

	switch chainType {
	case "ethereum":
		generator = blockchain.Ethereum{}
	default:
		err := c.String(http.StatusBadRequest, "Invalid chain type")
		return err
	}

	address := generator.GenerateAddress()
	return c.String(http.StatusOK, address)
}
