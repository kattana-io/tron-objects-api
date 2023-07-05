package url

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

type APIURLProvider interface {
	Request(url string, body []byte) ([]byte, error)
	GetBlockByNum() string
	GetTransactionInfoByID() string
	GetTransactionInfoByBlockNum() string
	TriggerConstantContract() string
	TriggerSmartContract() string
	GetContractInfo() string
}

type NodeURLProvider struct {
	host string
}

func (n *NodeURLProvider) TriggerSmartContract() string {
	return fmt.Sprintf("%s/wallet/triggersmartcontract", n.host)
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

func (n *NodeURLProvider) GetTransactionInfoByBlockNum() string {
	return fmt.Sprintf("%s/wallet/gettransactioninfobyblocknum", n.host)
}

const maxTimeout = 30 * time.Second

func (n *NodeURLProvider) Request(url string, body []byte) ([]byte, error) {
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
	defer res.Body.Close()

	resBody, err2 := io.ReadAll(res.Body)

	return resBody, err2
}

func NewNodeURLProvider(host string) APIURLProvider {
	return &NodeURLProvider{
		host: host,
	}
}
