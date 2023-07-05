package url

import (
	"bytes"
	"context"
	"net/http"
	"os"
)

/**
 * TronGrid URLs
 */
const trongridHost = "https://api.trongrid.io"

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
	ctx, cancel := context.WithTimeout(context.Background(), maxTimeout)
	defer cancel()

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	if n.APIKey != "" {
		req.Header.Add("TRON-PRO-API-KEY", n.APIKey)
	}
	req.Header.Add("Content-Type", "application/json")
	req = req.WithContext(ctx)
	client := http.DefaultClient
	resp, err = client.Do(req)

	return resp, err
}

func (n *TrongridURLProvider) GetTransactionInfoByBlockNum() string {
	return trongridHost + "/wallet/gettransactioninfobyblocknum"
}

func (n *TrongridURLProvider) GetBlockByNum() string {
	return trongridHost + "/wallet/getblockbynum"
}

func (n *TrongridURLProvider) GetTransactionInfoByID() string {
	return trongridHost + "/wallet/gettransactioninfobyid"
}

func (n *TrongridURLProvider) TriggerConstantContract() string {
	return trongridHost + "/wallet/triggerconstantcontract"
}

func (n *TrongridURLProvider) GetContractInfo() string {
	return trongridHost + "/wallet/getcontractinfo"
}

func (n *TrongridURLProvider) TriggerSmartContract() string {
	return trongridHost + "/wallet/triggersmartcontract"
}
