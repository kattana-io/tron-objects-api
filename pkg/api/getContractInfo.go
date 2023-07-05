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

func (a *API) GetContractInfo(token *Address) (ContractInfoResponse, error) {
	postBody, _ := json.Marshal(map[string]any{
		"value":   token.ToHex(),
		"visible": "false",
	})

	body, err := a.provider.Request(a.provider.GetContractInfo(), postBody)
	if err != nil {
		a.log.Error(err.Error())
		return ContractInfoResponse{}, err
	}

	var data ContractInfoResponse
	err2 := json.Unmarshal(body, &data)
	if err2 != nil {
		return ContractInfoResponse{}, err
	}

	return data, nil
}
