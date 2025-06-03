package rest

import (
	"encoding/json"
)

type GetTransactionInfoByBlockNumResp = []*GetTransactionInfoByBlockNumData

type GetTransactionInfoByBlockNumData struct {
	Log             []Log    `json:"log,omitempty"`
	Fee             int      `json:"fee,omitempty"`
	BlockNumber     int      `json:"blockNumber"`
	ContractResult  []string `json:"contractResult"`
	BlockTimeStamp  int64    `json:"blockTimeStamp"`
	Receipt         Receipt  `json:"receipt"`
	ID              string   `json:"id"`
	ContractAddress string   `json:"contract_address,omitempty"`
	Result          string   `json:"result,omitempty"`
	ResMessage      string   `json:"resMessage,omitempty"`
}

type Receipt struct {
	Result             string `json:"result"`
	NetFee             int    `json:"net_fee"`
	EnergyPenaltyTotal int    `json:"energy_penalty_total"`
	EnergyFee          int    `json:"energy_fee"`
	EnergyUsageTotal   int    `json:"energy_usage_total"`
	OriginEnergyUsage  int    `json:"origin_energy_usage"`
}

func (a *API) GetTransactionInfoByBlockNum(blockNumber int64) (GetTransactionInfoByBlockNumResp, error) {
	postBody, _ := json.Marshal(map[string]any{
		"num": blockNumber,
	})

	body, err := a.provider.Request(a.provider.GetTransactionInfoByBlockNum(), postBody)
	if err != nil {
		a.log.Sugar().Warnf("Could not load tx: %v", blockNumber)
		a.log.Error(err.Error())
		return GetTransactionInfoByBlockNumResp{}, err
	}

	var data GetTransactionInfoByBlockNumResp
	err2 := json.Unmarshal(body, &data)

	if err2 != nil {
		a.log.Sugar().Warnf("Could not load txs: %v", blockNumber)
		return GetTransactionInfoByBlockNumResp{}, err
	}

	return data, nil
}
