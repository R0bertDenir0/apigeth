package modules

import (
	"context"
	"fmt"
	"github.com/R0bertDenir0/apigeth/models"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func GetTheLatestBlock(client ethclient.Client) *models.Block {
	//recover function from panics to prevent API crashing
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	//Query the latest block
	header, _ := client.HeaderByNumber(context.Background(), nil)
	blockNumber := big.NewInt(header.Number.Int64())
	block, err := client.BlockByNumber(context.Background(), blockNumber)

	if err != nil {
		log.Fatalln(err)
	}

	//response to our model

	_block := &models.Block{
		BlockNumber:       block.Number().Int64(),
		TimeStamp:         block.Time(),
		Difficulty:        block.Difficulty().Uint64(),
		Hash:              block.Hash().String(),
		TransactionsCount: len(block.Transactions()),
		Transactions:      []models.Transaction{},
	}

	for _, tx := range block.Transactions() {
		_block.Transactions = append(_block.Transactions, models.Transaction{
			Hash:     tx.Hash().String(),
			Value:    tx.Value().String(),
			Gas:      tx.Gas(),
			GasPrice: tx.GasPrice().Uint64(),
			Nonce:    tx.Nonce(),
			To:       tx.To().String(),
		})

	}
	return _block

}
