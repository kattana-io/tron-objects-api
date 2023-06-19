package url

import (
	"bytes"
	"net/http"
	"os"
)

/**
 * TronGrid URLs
 */

func NewTrongridURLProvider() APIURLProvider {
	return &TrongridURLProvider{
		APIKey: os.Getenv("TRONGRID_API_KEY"),
	}
}

type TrongridURLProvider struct {
	APIKey string
}

// Request - Add headers to request https://developers.tron.network/reference/api-key#how-to-use-api-keys
func (n *TrongridURLProvider) Request(url string, body []byte) (resp *http.Response, err error) {
	client := &http.Client{}
	//nolint:noctx
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	if n.APIKey != "" {
		req.Header.Add("TRON-PRO-API-KEY", n.APIKey)
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err = client.Do(req)

	return resp, err
}

func (n *TrongridURLProvider) GetBlockByNum() string {
	return "https://api.trongrid.io/wallet/getblockbynum"
}

func (n *TrongridURLProvider) GetTransactionInfoByID() string {
	return "https://api.trongrid.io/wallet/gettransactioninfobyid"
}

func (n *TrongridURLProvider) TriggerConstantContract() string {
	return "https://api.trongrid.io/wallet/triggerconstantcontract"
}

func (n *TrongridURLProvider) GetContractInfo() string {
	return "https://api.trongrid.io/wallet/getcontractinfo"
}
