package rest

import "encoding/json"

type GetBlockByNumResp struct {
	BlockID     string `json:"blockID"`
	BlockHeader struct {
		RawData struct {
			Number         int    `json:"number"`
			TxTrieRoot     string `json:"txTrieRoot"`
			WitnessAddress string `json:"witness_address"`
			ParentHash     string `json:"parentHash"`
			Version        int    `json:"version"`
			Timestamp      int64  `json:"timestamp"`
		} `json:"raw_data"`
		WitnessSignature string `json:"witness_signature"`
	} `json:"block_header"`
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Ret []struct {
		ContractRet string `json:"contractRet"`
	} `json:"ret"`
	Signature []string `json:"signature"`
	TxID      string   `json:"txID"`
	RawData   struct {
		Contract []struct {
			Parameter struct {
				Value struct {
					Amount          int64  `json:"amount,omitempty"`
					AssetName       string `json:"asset_name,omitempty"`
					OwnerAddress    string `json:"owner_address"`
					ToAddress       string `json:"to_address,omitempty"`
					Data            string `json:"data,omitempty"`
					ContractAddress string `json:"contract_address,omitempty"`
					Resource        string `json:"resource,omitempty"`
					ReceiverAddress string `json:"receiver_address,omitempty"`
					FrozenDuration  int    `json:"frozen_duration,omitempty"`
					FrozenBalance   int    `json:"frozen_balance,omitempty"`
					CallValue       int    `json:"call_value,omitempty"`
				} `json:"value"`
				TypeURL string `json:"type_url"`
			} `json:"parameter"`
			Type string `json:"type"`
		} `json:"contract"`
		RefBlockBytes string `json:"ref_block_bytes"`
		RefBlockHash  string `json:"ref_block_hash"`
		Expiration    int64  `json:"expiration"`
		FeeLimit      int    `json:"fee_limit,omitempty"`
		Timestamp     int64  `json:"timestamp"`
	} `json:"raw_data"`
	RawDataHex string `json:"raw_data_hex"`
}

func (a *API) GetBlockByNum(number int32) (*GetBlockByNumResp, error) {
	postBody, _ := json.Marshal(map[string]any{
		"num": number,
	})

	body, err := a.provider.Request(a.provider.GetBlockByNum(), postBody)
	if err != nil {
		a.log.Error(err.Error())
		return &GetBlockByNumResp{}, err
	}

	var data GetBlockByNumResp
	err2 := json.Unmarshal(body, &data)

	if err2 != nil {
		return &GetBlockByNumResp{}, err
	}

	return &data, nil
}
