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
func (t *EwalletContract) CreateTrans(
    ctx contractapi.TransactionContextInterface, 
    assetID  string, createDate time.Time, tranHash string) error {
    exists, err := t.AssetExists(ctx, assetID)
    if err != nil {
        return fmt.Errorf("Failed to get Asset: %v", err)
    }
    if exists {
        return fmt.Errorf("Asset already exists: %s", assetID)
    }
    asset := &Asset{
        Id:             assetID,
        CreateDate:     createDate,
        TranHash:       tranHash,
    }
    assetBytes, err := json.Marshal(asset)
    if err != nil {
        return err
    }
    err = ctx.GetStub().PutState(assetID, assetBytes)
    if err != nil { return err }
    // Registering the event
    eventPayload := fmt.Sprintf("Asset with ID %s created", assetID)
    err = ctx.GetStub().SetEvent("CreateTransEvent", []byte(eventPayload))
    if err != nil { return fmt.Errorf("Failed to set event: %v", err)}
    return nil
}
