package models

import "github.com/ethereum/go-ethereum/common"

type Block struct {
	BaseFeePerGas    string        `json:"baseFeePerGas"`
	Difficulty       string        `json:"difficulty"`
	ExtraData        string        `json:"extraData"`
	GasLimit         string        `json:"gasLimit"`
	GasUsed          string        `json:"gasUsed"`
	Hash             string        `json:"hash"`
	LogsBloom        string        `json:"logsBloom"`
	Miner            string        `json:"miner"`
	MixHash          string        `json:"mixHash"`
	Nonce            string        `json:"nonce"`
	Number           string        `json:"number"`
	ParentHash       string        `json:"parentHash"`
	ReceiptsRoot     string        `json:"receiptsRoot"`
	Sha3Uncles       string        `json:"sha3Uncles"`
	Size             string        `json:"size"`
	StateRoot        string        `json:"stateRoot"`
	Timestamp        string        `json:"timestamp"`
	TotalDifficulty  string        `json:"totalDifficulty"`
	Transactions     []Transaction `json:"transactions"`
	TransactionsRoot string        `json:"transactionsRoot"`
	Uncles           []any         `json:"uncles"`
}

type Transaction struct {
	BlockHash        string  `json:"blockHash"`
	BlockNumber      string  `json:"blockNumber"`
	From             string  `json:"from"`
	Gas              string  `json:"gas"`
	GasPrice         string  `json:"gasPrice"`
	Hash             string  `json:"hash"`
	Input            string  `json:"input"`
	Nonce            string  `json:"nonce"`
	R                string  `json:"r"`
	S                string  `json:"s"`
	To               *string `json:"to"`
	TransactionIndex string  `json:"transactionIndex"`
	Type             string  `json:"type"`
	V                string  `json:"v"`
	Value            string  `json:"value"`
}

type ContractInfo struct {
	RuntimeCode string `json:"runtimecode"`
	// TODO check how to add smart contract info
	// SmartContract struct {
	// 	ConsumeUserResourcePercent int    `json:"consume_user_resource_percent"`
	// 	OriginAddress              string `json:"origin_address"`
	// 	ContractAddress            string `json:"contract_address"`
	// 	CodeHash                   string `json:"code_hash"`
	// } `json:"smart_contract"`
}

type TransactionReceipt struct {
	BlockHash         string  `json:"blockHash"`
	BlockNumber       string  `json:"blockNumber"`
	ContractAddress   *string `json:"contractAddress"`
	CumulativeGasUsed string  `json:"cumulativeGasUsed"`
	EffectiveGasPrice string  `json:"effectiveGasPrice"`
	From              string  `json:"from"`
	GasUsed           string  `json:"gasUsed"`
	Logs              []Log   `json:"logs"`
	LogsBloom         string  `json:"logsBloom"`
	Status            string  `json:"status"`
	To                string  `json:"to"`
	TransactionHash   string  `json:"transactionHash"`
	TransactionIndex  string  `json:"transactionIndex"`
	Type              string  `json:"type"`
}

type Log struct {
	Address          string   `json:"address"`
	BlockHash        string   `json:"blockHash"`
	BlockNumber      string   `json:"blockNumber"`
	Data             string   `json:"data"`
	LogIndex         string   `json:"logIndex"`
	Removed          bool     `json:"removed"`
	Topics           []string `json:"topics"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
}

type GetLogsRequest struct {
	FromBlock string           `json:"fromBlock,omitempty"`
	ToBlock   string           `json:"toBlock,omitempty"`
	Address   []common.Address `json:"address,omitempty"`
	Topics    []common.Hash    `json:"topics,omitempty"`
	BlockHash common.Hash      `json:"blockHash,omitempty"`
}
