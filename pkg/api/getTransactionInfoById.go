package api

import (
	"encoding/json"
	"errors"
)

type GetTransactionInfoByIdResp struct {
	Error string `json:"Error"`

	Id              string   `json:"id"`
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

func (a *Api) GetTransactionInfoById(id string) (*GetTransactionInfoByIdResp, error) {
	postBody, _ := json.Marshal(map[string]interface{}{
		"value": id,
	})

	res, err := a.provider.Request(a.provider.GetTransactionInfoById(), postBody)
	if err != nil {
		a.log.Warn("Could not load tx: " + id)
		a.log.Error(err.Error())
		return &GetTransactionInfoByIdResp{}, err
	}

	defer res.Body.Close()

	var data GetTransactionInfoByIdResp
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	if err != nil {
		a.log.Warn("Could not load tx: " + id)
		return &GetTransactionInfoByIdResp{}, err
	}

	if data.Error != "" {
		return nil, errors.New(" got error during tx call: " + data.Error)
	}

	return &data, nil
}
