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
	// Create acc
	//src, err := account.FromPrivateKeyHex(srcPrivKey)
	//if err != nil {
	//	log.Fatal("Failed to parse private key hex - ", err)
	//}
	impl := api.NewAPI("", nil, url.NewTrongridURLProvider())
	r := sunswap.NewRouter(impl)
	//
	bts, selector, err := r.SwapETHToToken(*api.FromBase58(usdt), big.NewInt(10^18), *api.FromBase58(ZeroAddress), big.NewInt(time.Now().Add(time.Minute).Unix()))
	if err != nil {
		log.Fatal(err)
	}
	tx, err := impl.TriggerSmartContract(r.ContractAddress(), selector, hexutil.Encode(bts)[2:], big.NewInt(0), big.NewInt(0))
	fmt.Println(tx)
}
