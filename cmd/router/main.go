package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/kattana-io/tron-objects-api/pkg/api"
	"github.com/kattana-io/tron-objects-api/pkg/sunswap"
	"github.com/kattana-io/tron-objects-api/pkg/url"
	"log"
	"math/big"
	"time"
)

const (
	srcPrivKey  = "000000000000000000000000000000000000000000000000000000000000010f"
	usdt        = "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
	ZeroAddress = "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
)

func main() {
	impl := api.NewAPI("", nil, url.NewTrongridURLProvider())
	r := sunswap.NewRouter(impl)
	//nolint:gomnd
	OneTrx, _ := big.NewInt(0).SetString("1000000000000000000", 10)
	//nolint:lll
	bts, selector, err := r.SwapETHToTokens(*api.FromBase58(usdt), OneTrx, *api.FromBase58(ZeroAddress), big.NewInt(time.Now().Add(time.Minute).Unix()))
	if err != nil {
		log.Fatal(err)
	}
	tx, _ := impl.TriggerSmartContract(r.ContractAddress(), selector, hexutil.Encode(bts)[2:], big.NewInt(0), big.NewInt(0))
	fmt.Println(tx)
}
