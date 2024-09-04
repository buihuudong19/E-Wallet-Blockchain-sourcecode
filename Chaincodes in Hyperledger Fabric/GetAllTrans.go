package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type EwalletContract struct {
	contractapi.Contract
}

func (t *EwalletContract) GetAllTrans(
	ctx contractapi.TransactionContextInterface) ([]Transaction, error) {
	queryString := `{
						"selector": {
							"ID": {
								"$gt": null
							}
						}
					}`
	transactions, err := t.getQueryResultForQueryString(ctx, queryString)

	if err != nil {
		return nil, fmt.Errorf("Failed to get transactions by conditions %v", err)
	}
	return transactions, nil
}
