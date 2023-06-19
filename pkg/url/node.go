package url

import (
	"bytes"
	"fmt"
	"net/http"
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

func (n *NodeURLProvider) Request(url string, body []byte) (*http.Response, error) {
	//nolint:gosec
	res, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	return res, err
}

func NewNodeURLProvider(host string) APIURLProvider {
	return &NodeURLProvider{
		host: host,
	}
}
