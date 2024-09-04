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

func (s *EwalletContract) UpdateTranById(
    ctx contractapi.TransactionContextInterface, 
    assetID string,  
    createDate time.Time,
    tranHash string) error {
    exists, err := s.AssetExists(ctx, id)
    if err != nil { return err }
    if !exists {
        return fmt.Errorf("the asset %s does not exist", id)
    }
    asset := &Asset{
        Id:             assetID,
        CreateDate:     createDate,
        TranHash:       tranHash,
    }
    assetJSON, err := json.Marshal(asset)
    if err != nil { return err}
    err = ctx.GetStub().PutState(id, assetJSON)
    if err != nil { return err }
    // Registering the event when update ledger
    eventPayload := fmt.Sprintf("Asset with ID %s updated", assetID)
    err = ctx.GetStub().SetEvent("UpdateTransEvent", []byte(eventPayload))
    if err != nil { return fmt.Errorf("Failed to set event: %v", err)}
    return nil
}
