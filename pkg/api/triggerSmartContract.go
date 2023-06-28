package api

import (
	"github.com/goccy/go-json"
	"math/big"
)

type TSMResponse struct {
	Result struct {
		Result bool `json:"result"`
	} `json:"result"`
	Transaction struct {
		Visible bool   `json:"visible"`
		TxID    string `json:"txID"`
		RawData struct {
			Contract []struct {
				Parameter struct {
					Value struct {
						Data            string `json:"data"`
						OwnerAddress    string `json:"owner_address"`
						ContractAddress string `json:"contract_address"`
					} `json:"value"`
					TypeURL string `json:"type_url"`
				} `json:"parameter"`
				Type string `json:"type"`
			} `json:"contract"`
			RefBlockBytes string `json:"ref_block_bytes"`
			RefBlockHash  string `json:"ref_block_hash"`
			Expiration    int64  `json:"expiration"`
			Timestamp     int64  `json:"timestamp"`
		} `json:"raw_data"`
		RawDataHex string `json:"raw_data_hex"`
	} `json:"transaction"`
}

// TriggerSmartContract - returns an unsigned transaction /**
func (a *API) TriggerSmartContract(contract, selector, parameter string, feeLimit, callValue *big.Int) (*TSMResponse, error) {
	input := map[string]any{
		"owner_address":     contract,
		"contract_address":  contract,
		"function_selector": selector,
		"parameter":         parameter,
		"fee_limit":         feeLimit,
		"call_value":        callValue,
		"visible":           true,
	}

	postBody, err := json.Marshal(&input)
	if err != nil {
		return nil, err
	}

	var result TSMResponse

	res, err := a.provider.Request(a.provider.TriggerSmartContract(), postBody)
	if err != nil {
		a.log.Error(err.Error())
		return &result, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&result)
	if err != nil {
		return &result, err
	}

	return &result, nil
}
