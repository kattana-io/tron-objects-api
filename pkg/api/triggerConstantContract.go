package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"net/http"
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

// TriggerConstantContract /**
func (a *API) TriggerConstantContract(contractAddress, functionSelector, parameter string) (*TCCResponse, error) {
	postBody, _ := json.Marshal(map[string]any{
		"owner_address":     DummyCaller,
		"contract_address":  contractAddress,
		"function_selector": functionSelector,
		"parameter":         parameter,
	})

	res, err := a.provider.Request(a.provider.TriggerConstantContract(), postBody)
	if err != nil {
		a.log.Error(err.Error())
		return &TCCResponse{}, err
	}

	defer res.Body.Close()

	var data TCCResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return &TCCResponse{}, err
	}

	return &data, nil
}

const defaultDecimals = 18

func (a *API) GetTokenDecimals(token string) (int32, error) {
	postBody, _ := json.Marshal(map[string]any{
		"owner_address":     DummyCaller,
		"contract_address":  token,
		"function_selector": "decimals()",
	})

	res, err := a.provider.Request(a.provider.TriggerConstantContract(), postBody)
	if err != nil {
		a.log.Error(err.Error())
		return defaultDecimals, err
	}

	defer res.Body.Close()

	var data TCCResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	if err != nil || len(data.ConstantResult) == 0 {
		return defaultDecimals, err
	}

	result, err := strconv.ParseInt(TrimZeroes(data.ConstantResult[0]), 16, 16)
	if err != nil {
		return defaultDecimals, err
	}
	return int32(result), nil
}

//nolint:dupl
func (a *API) GetPairToken(pair string) (string, error) {
	postBody, _ := json.Marshal(map[string]interface{}{
		"owner_address":     DummyCaller,
		"contract_address":  pair,
		"function_selector": "tokenAddress()",
	})

	res, err := a.provider.Request(a.provider.TriggerConstantContract(), postBody)
	if err != nil {
		a.log.Error(err.Error())
		return "", err
	}
	defer res.Body.Close()

	var data TCCResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	if err != nil || len(data.ConstantResult) == 0 {
		return "", err
	}

	return TrimZeroes(data.ConstantResult[0]), nil
}

//nolint:dupl
func (a *API) GetToken0(pair string) (string, error) {
	postBody, _ := json.Marshal(map[string]interface{}{
		"owner_address":     DummyCaller,
		"contract_address":  pair,
		"function_selector": "token0()",
	})

	res, err := a.provider.Request(a.provider.TriggerConstantContract(), postBody)
	if err != nil {
		a.log.Error(err.Error())
		return "", err
	}

	defer res.Body.Close()

	var data TCCResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	if err != nil || len(data.ConstantResult) == 0 {
		return "", err
	}

	return TrimZeroes(data.ConstantResult[0]), nil
}

//nolint:dupl
func (a *API) GetToken1(pair string) (string, error) {
	postBody, _ := json.Marshal(map[string]interface{}{
		"owner_address":     DummyCaller,
		"contract_address":  pair,
		"function_selector": "token1()",
	})

	res, err := a.provider.Request(a.provider.TriggerConstantContract(), postBody)
	if err != nil {
		a.log.Error(err.Error())
		return "", err
	}
	defer res.Body.Close()

	var data TCCResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	if err != nil || len(data.ConstantResult) == 0 {
		return "", err
	}

	return TrimZeroes(data.ConstantResult[0]), nil
}

func (a *API) GetTokenName(hexAddress string) (string, error) {
	postBody, _ := json.Marshal(map[string]interface{}{
		"owner_address":     DummyCaller,
		"contract_address":  hexAddress,
		"function_selector": "name()",
	})
	responseBody := bytes.NewBuffer(postBody)

	res, err := http.Post(a.provider.TriggerConstantContract(), "application/json", responseBody)
	if err != nil {
		a.log.Error(err.Error())
		return "", err
	}
	defer res.Body.Close()

	var data TCCResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	if err != nil || len(data.ConstantResult) == 0 {
		return "", err
	}

	decodedData, err := hexutil.Decode(fmt.Sprintf("0x%s", data.ConstantResult[0][128:]))
	decodedData = bytes.TrimRight(decodedData, "\u0000")
	if err != nil {
		a.log.Error(err.Error())
		return "", err
	}

	return string(decodedData), nil
}

func (a *API) GetTokenSymbol(hexAddress string) (string, error) {
	postBody, _ := json.Marshal(map[string]interface{}{
		"owner_address":     DummyCaller,
		"contract_address":  hexAddress,
		"function_selector": "symbol()",
	})
	responseBody := bytes.NewBuffer(postBody)

	//nolint:noctx
	res, err := http.Post(a.provider.TriggerConstantContract(), "application/json", responseBody)
	if err != nil {
		a.log.Error(err.Error())
		return "", err
	}
	defer res.Body.Close()

	var data TCCResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
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
