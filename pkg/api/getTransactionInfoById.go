package api

import (
	"encoding/json"
	"errors"
)

type GetTransactionInfoByIDResp struct {
	Error string `json:"Error"`

	ID              string   `json:"id"`
	Fee             int      `json:"fee"`
	BlockNumber     int      `json:"blockNumber"`
	BlockTimeStamp  int64    `json:"blockTimeStamp"`
	ContractResult  []string `json:"contractResult"`
	ContractAddress string   `json:"contract_address"`
	Receipt         struct {
		OriginEnergyUsage int    `json:"origin_energy_usage"`
		EnergyUsageTotal  int    `json:"energy_usage_total"`
		NetFee            int    `json:"net_fee"`
		Result            string `json:"result"`
	} `json:"receipt"`
	Log []Log `json:"log"`
}

type Log struct {
	Address string   `json:"address"`
	Topics  []string `json:"topics"`
	Data    string   `json:"data"`
}

func (a *API) GetTransactionInfoByID(id string) (*GetTransactionInfoByIDResp, error) {
	postBody, _ := json.Marshal(map[string]any{
		"value": id,
	})

	res, err := a.provider.Request(a.provider.GetTransactionInfoByID(), postBody)
	if err != nil {
		a.log.Warn("Could not load tx: " + id)
		a.log.Error(err.Error())
		return &GetTransactionInfoByIDResp{}, err
	}

	defer res.Body.Close()

	var data GetTransactionInfoByIDResp
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	if err != nil {
		a.log.Warn("Could not load tx: " + id)
		return &GetTransactionInfoByIDResp{}, err
	}

	if data.Error != "" {
		return nil, errors.New(" got error during tx call: " + data.Error)
	}

	return &data, nil
}
