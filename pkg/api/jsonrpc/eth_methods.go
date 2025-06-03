package jsonrpc

import (
	"context"
	"errors"
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

func (a *API) GetContractInfo(ctx context.Context, contract common.Address) (*models.ContractInfo, error) {
	runtimeCode, err := a.ethclient.CodeAt(ctx, contract, nil)
	if err != nil {
		a.log.Error("error getting runtime code", zap.String("contract", contract.String()), zap.String("error", err.Error()))
		return nil, err
	}
	// TODO add another smart contract info if it is possible
	return &models.ContractInfo{RuntimeCode: string(runtimeCode)}, nil
}
