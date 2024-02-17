package tests

import (
	"chain/api/handlers"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type addressResponse struct {
	PrivateKey string `json:"private_key"`
	Address    string `json:"address"`
}

func TestEthereumAddress(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/?type=ethereum", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handlers.GenerateAddress(c)) {
		var resp addressResponse
		err := json.Unmarshal(rec.Body.Bytes(), &resp)
		assert.NoError(t, err, "JSON response should be correctly parsed")

		// 从私钥的十六进制字符串解码
		privateKeyBytes, err := hex.DecodeString(resp.PrivateKey)
		assert.NoError(t, err, "Should decode private key without error")

		// 从私钥字节恢复私钥
		privateKey, err := crypto.ToECDSA(privateKeyBytes)
		assert.NoError(t, err, "Should recover private key without error")

		// 从私钥生成公钥，再生成以太坊地址验证
		derivedAddress := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()

		// 检查解码的地址是否与派生地址匹配
		assert.Equal(t, resp.Address, derivedAddress, "Derived address should match the provided address")
	}
}
