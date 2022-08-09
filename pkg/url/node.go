package url

import (
	"bytes"
	"fmt"
	"net/http"
)

type ApiUrlProvider interface {
	Request(url string, body []byte) (*http.Response, error)
	GetBlockByNum() string
	GetTransactionInfoById() string
	TriggerConstantContract() string
}

type NodeUrlProvider struct {
	host    string
	network string
}

func (n *NodeUrlProvider) GetBlockByNum() string {
	return fmt.Sprintf("%s/wallet/getblockbynum", n.network)
}

func (n *NodeUrlProvider) GetTransactionInfoById() string {
	return fmt.Sprintf("%s/walletsolidity/gettransactioninfobyid", n.host)
}

func (n *NodeUrlProvider) TriggerConstantContract() string {
	return fmt.Sprintf("%s/walletsolidity/triggerconstantcontract", n.host)
}

func (n *NodeUrlProvider) Request(url string, body []byte) (*http.Response, error) {
	res, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	return res, err
}

func NewNodeUrlProvider(host string, network string) ApiUrlProvider {
	return &NodeUrlProvider{
		host:    host,
		network: network,
	}
}
