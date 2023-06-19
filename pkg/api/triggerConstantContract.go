package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"strconv"
)

const DummyCaller = "410000000000000000000000000000000000000000"

type TCCResponse struct {
	Result struct {
		Result bool `json:"result"`
	} `json:"result"`
	EnergyUsed     int      `json:"energy_used"`
	ConstantResult []string `json:"constant_result"`
	Transaction    struct {
		Ret []struct {
		} `json:"ret"`
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

func (a *API) TCCRequest(input map[string]any) (*TCCResponse, error) {
	postBody, _ := json.Marshal(input)
	var result TCCResponse

	res, err := a.provider.Request(a.provider.TriggerConstantContract(), postBody)
	if err != nil {
		a.log.Error(err.Error())
		return &TCCResponse{}, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&result)
	if err != nil {
		return &result, err
	}

	return &result, nil
}

// TriggerConstantContract /**
func (a *API) TriggerConstantContract(contractAddress, functionSelector, parameter string) (*TCCResponse, error) {
	return a.TCCRequest(map[string]any{
		"owner_address":     DummyCaller,
		"contract_address":  contractAddress,
		"function_selector": functionSelector,
		"parameter":         parameter,
	})
}

const defaultDecimals = 18

func (a *API) GetTokenDecimals(token string) (int32, error) {
	data, err := a.TCCRequest(map[string]any{
		"owner_address":     DummyCaller,
		"contract_address":  token,
		"function_selector": "decimals()",
	})

	if err != nil || len(data.ConstantResult) == 0 {
		return defaultDecimals, err
	}

	result, err := strconv.ParseInt(TrimZeroes(data.ConstantResult[0]), 16, 16)
	if err != nil {
		return defaultDecimals, err
	}
	return int32(result), nil
}

func (a *API) GetPairToken(pair string) (string, error) {
	data, err := a.TCCRequest(map[string]any{
		"owner_address":     DummyCaller,
		"contract_address":  pair,
		"function_selector": "tokenAddress()",
	})

	if err != nil || len(data.ConstantResult) == 0 {
		return "", err
	}

	return TrimZeroes(data.ConstantResult[0]), nil
}

func (a *API) GetToken0(pair string) (string, error) {
	data, err := a.TCCRequest(map[string]any{
		"owner_address":     DummyCaller,
		"contract_address":  pair,
		"function_selector": "token0()",
	})

	if err != nil || len(data.ConstantResult) == 0 {
		return "", err
	}

	return TrimZeroes(data.ConstantResult[0]), nil
}

func (a *API) GetToken1(pair string) (string, error) {
	data, err := a.TCCRequest(map[string]any{
		"owner_address":     DummyCaller,
		"contract_address":  pair,
		"function_selector": "token1()",
	})

	if err != nil || len(data.ConstantResult) == 0 {
		return "", err
	}

	return TrimZeroes(data.ConstantResult[0]), nil
}

func (a *API) GetTokenName(hexAddress string) (string, error) {
	data, err := a.TCCRequest(map[string]any{
		"owner_address":     DummyCaller,
		"contract_address":  hexAddress,
		"function_selector": "name()",
	})

	if err != nil || len(data.ConstantResult) == 0 {
		return "", err
	}

	decodedData, err := hexutil.Decode(fmt.Sprintf("0x%s", data.ConstantResult[0][128:]))
	if err != nil {
		a.log.Error(err.Error())
		return "", err
	}
	decodedData = bytes.TrimRight(decodedData, "\u0000")

	return string(decodedData), nil
}

func (a *API) GetTokenSymbol(hexAddress string) (string, error) {
	data, err := a.TCCRequest(map[string]any{
		"owner_address":     DummyCaller,
		"contract_address":  hexAddress,
		"function_selector": "symbol()",
	})

	if err != nil || len(data.ConstantResult) == 0 {
		return "", err
	}

	decodedData, err := hexutil.Decode(fmt.Sprintf("0x%s", data.ConstantResult[0][128:]))
	if err != nil {
		return "", err
	}
	decodedData = bytes.TrimRight(decodedData, "\u0000")

	return string(decodedData), nil
}
