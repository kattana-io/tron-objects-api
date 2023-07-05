package api

import (
	"encoding/json"
)

type GetTransactionInfoByBlockNumResp = []GetTransactionInfoByBlockNumData

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

func (a *API) GetTransactionInfoByBlockNum(id string) (*GetTransactionInfoByBlockNumResp, error) {
	postBody, _ := json.Marshal(map[string]any{
		"value": id,
	})

	res, err := a.provider.Request(a.provider.GetTransactionInfoByBlockNum(), postBody)
	if err != nil {
		a.log.Warn("Could not load tx: " + id)
		a.log.Error(err.Error())
		return &GetTransactionInfoByBlockNumResp{}, err
	}

	defer res.Body.Close()

	var data GetTransactionInfoByBlockNumResp
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	if err != nil {
		a.log.Warn("Could not load txs: " + id)
		return &GetTransactionInfoByBlockNumResp{}, err
	}

	return &data, nil
}
