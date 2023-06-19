package url

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"
)

type APIURLProvider interface {
	Request(url string, body []byte) (*http.Response, error)
	GetBlockByNum() string
	GetTransactionInfoByID() string
	TriggerConstantContract() string
	GetContractInfo() string
}

type NodeURLProvider struct {
	host string
}

func (n *NodeURLProvider) GetContractInfo() string {
	return fmt.Sprintf("%s/wallet/getcontractinfo", n.host)
}

func (n *NodeURLProvider) GetBlockByNum() string {
	return fmt.Sprintf("%s/wallet/getblockbynum", n.host)
}

func (n *NodeURLProvider) GetTransactionInfoByID() string {
	return fmt.Sprintf("%s/wallet/gettransactioninfobyid", n.host)
}

func (n *NodeURLProvider) TriggerConstantContract() string {
	return fmt.Sprintf("%s/wallet/triggerconstantcontract", n.host)
}

const maxTimeout = 30 * time.Second

func (n *NodeURLProvider) Request(url string, body []byte) (*http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), maxTimeout)
	defer cancel()

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func NewNodeURLProvider(host string) APIURLProvider {
	return &NodeURLProvider{
		host: host,
	}
}
