package url

import (
	"bytes"
	"net/http"
	"os"
)

/**
 * TronGrid URLs
 */

func NewTrongridUrlProvider() ApiUrlProvider {
	return &TrongridUrlProvider{
		ApiKey: os.Getenv("TRONGRID_API_KEY"),
	}
}

type TrongridUrlProvider struct {
	ApiKey string
}

// Request - Add headers to request https://developers.tron.network/reference/api-key#how-to-use-api-keys
func (n *TrongridUrlProvider) Request(url string, body []byte) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if n.ApiKey != "" {
		req.Header.Add("TRON-PRO-API-KEY", n.ApiKey)
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)

	return resp, err
}

func (n *TrongridUrlProvider) GetBlockByNum() string {
	return "https://api.trongrid.io/wallet/getblockbynum"
}

func (n *TrongridUrlProvider) GetTransactionInfoById() string {
	return "https://api.trongrid.io/wallet/gettransactioninfobyid"
}

func (n *TrongridUrlProvider) TriggerConstantContract() string {
	return "https://api.trongrid.io/wallet/triggerconstantcontract"
}

func (n *TrongridUrlProvider) GetContractInfo() string {
	return "https://api.trongrid.io/wallet/getcontractinfo"
}
