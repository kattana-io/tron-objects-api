package jsonrpc

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/kattana-io/tron-objects-api/pkg/models"
	"go.uber.org/zap"
)

func (a *API) GetBlockNumber(ctx context.Context) (uint64, error) {
	num, err := a.ethclient.BlockNumber(ctx)
	if err != nil {
		a.log.Error("error getting last block number", zap.String("error", err.Error()))
		return 0, err
	}
	return num, nil
}

func (a *API) GetSuggestGasPrice(ctx context.Context) (*big.Int, error) {
	gas, err := a.ethclient.SuggestGasPrice(ctx)
	if err != nil {
		a.log.Error("error getting suggest gas price", zap.String("error", err.Error()))
		return nil, err
	}
	return gas, nil
}

func (a *API) GetBalanceAt(ctx context.Context, account common.Address, block *big.Int) (*big.Int, error) {
	gas, err := a.ethclient.BalanceAt(ctx, account, block)
	if err != nil {
		a.log.Error("error getting balance", zap.String("error", err.Error()))
		return nil, err
	}
	return gas, nil
}

func (a *API) GetContractInfo(ctx context.Context, contract common.Address) (*models.ContractInfo, error) {
	runtimeCode, err := a.ethclient.CodeAt(ctx, contract, nil)
	if err != nil {
		a.log.Error("error getting runtime code", zap.String("contract", contract.String()), zap.String("error", err.Error()))
		return nil, err
	}
	// TODO add another smart contract info if it is possible
	return &models.ContractInfo{RuntimeCode: string(runtimeCode)}, nil
}

func (a *API) GetBlockByNum(ctx context.Context, num int64) (*models.Block, error) {
	block, err := a.rpcclient.GetBlockByNum(ctx, num, true)
	if err != nil {
		a.log.Error("error getting block by num", zap.Int64("num", num), zap.String("error", err.Error()))
		return nil, err
	}
	return block, nil
}

func (a *API) GetTransactionByHash(ctx context.Context, hash common.Hash) (*models.Transaction, error) {
	if hash.String() == "" {
		return nil, errors.New("hash is empty")
	}

	tx, err := a.rpcclient.GetTransactionByHash(ctx, hash)
	if err != nil {
		a.log.Error("error getting tx by hash", zap.String("hash", hash.String()), zap.String("error", err.Error()))
		return nil, err
	}
	return tx, nil
}

func (a *API) GetTransactionByBlockNum(ctx context.Context, num int64) (*models.Transaction, error) {
	tx, err := a.rpcclient.GetTransactionByBlockNumAndIndex(ctx, num, 0)
	if err != nil {
		a.log.Error("error getting tx by block num", zap.Int64("num", num), zap.String("error", err.Error()))
		return nil, err
	}
	return tx, nil
}

func (a *API) GetTransactionReceipt(ctx context.Context, hash common.Hash) (*models.TransactionReceipt, error) {
	txReceipt, err := a.rpcclient.GetTransactionReceipt(ctx, hash)
	if err != nil {
		a.log.Error("error getting tx receipt", zap.String("hash", hash.String()), zap.String("error", err.Error()))
		return nil, err
	}
	return txReceipt, nil
}

func (a *API) GetLogs(ctx context.Context, req *models.GetLogsRequest) ([]models.Log, error) {
	logs, err := a.rpcclient.GetLogs(ctx, req)
	if err != nil {
		a.log.Error("error getting logs", zap.Any("req", req), zap.String("error", err.Error()))
		return nil, err
	}
	return logs, nil
}
