package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/kattana-io/tron-objects-api/pkg/api/rest"
	rest2 "github.com/kattana-io/tron-objects-api/pkg/client/rest"
	rest3 "github.com/kattana-io/tron-objects-api/pkg/connector/sunswap/rest"
	"github.com/kattana-io/tron-objects-api/pkg/types"
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
	impl := rest.NewAPI("", nil, rest2.NewTrongridURLProvider())
	r := rest3.NewRouter(impl)
	//nolint:gomnd
	OneTrx, _ := big.NewInt(0).SetString("1000000000000000000", 10) //nolint:mnd
	//nolint:lll
	bts, selector, err := r.SwapETHToTokens(*types.NewFromBase58(usdt), OneTrx, *types.NewFromBase58(ZeroAddress), big.NewInt(time.Now().Add(time.Minute).Unix()))
	if err != nil {
		log.Fatal(err)
	}
	tx, _ := impl.TriggerSmartContract(r.ContractAddress(), selector, hexutil.Encode(bts)[2:], big.NewInt(0), big.NewInt(0))
	fmt.Println(tx)
}
