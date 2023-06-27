package main

import (
	"encoding/json"
	"fmt"
	"github.com/kattana-io/go-tron"
	"github.com/kattana-io/go-tron/account"
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
	src, err := account.FromPrivateKeyHex(srcPrivKey)
	if err != nil {
		log.Fatal("Failed to parse private key hex - ", err)
	}
	impl := api.NewAPI("", nil, url.NewTrongridURLProvider())
	r := sunswap.NewRouter(impl)
	//
	bts, err := r.SwapETHToToken(*api.FromBase58(usdt), big.NewInt(10^18), *api.FromBase58(ZeroAddress), big.NewInt(time.Now().Add(time.Minute).Unix()))
	if err != nil {
		log.Fatal(err)
	}
	// create transaction
	RawData := json.RawMessage(bts)
	tx := tron.Transaction{
		Id:              "",
		Signatures:      nil,
		Results:         nil,
		ConstantResults: nil,
		Visible:         nil,
		RawData:         &RawData,
		RawDataHex:      nil,
		ContractAddress: nil,
	}
	err = src.Sign(&tx)
	if err != nil {
		log.Fatal(err)
	}
	// call trongrid to simulate tx
	fmt.Println(tx)
}
