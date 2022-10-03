package api

import (
	"encoding/json"
)

type ContractInfoResponse struct {
	Runtimecode   string `json:"runtimecode"`
	SmartContract struct {
		ConsumeUserResourcePercent int    `json:"consume_user_resource_percent"`
		OriginAddress              string `json:"origin_address"`
		ContractAddress            string `json:"contract_address"`
		CodeHash                   string `json:"code_hash"`
	} `json:"smart_contract"`
}

func (a *Api) GetContractInfo(token *Address) (ContractInfoResponse, error) {
	postBody, _ := json.Marshal(map[string]interface{}{
		"value":   token.ToHex(),
		"visible": "false",
	})

	res, err := a.provider.Request(a.provider.GetContractInfo(), postBody)
	if err != nil {
		a.log.Error(err.Error())
		return ContractInfoResponse{}, err
	}

	defer res.Body.Close()

	var data ContractInfoResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return ContractInfoResponse{}, err
	}

	return data, nil
}
