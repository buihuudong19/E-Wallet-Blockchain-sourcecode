func (t *EwalletContract) GetTranById(
    ctx contractapi.TransactionContextInterface, assetID string,
    createDate time.Time ) (*Asset, error) {
    queryString := fmt.Sprintf(`{
        "selector": {
            "Id": "%s",
            "CreateDate": "%s"
        }
    }`, assetType, status)
    assetBytes, err := ctx.GetStub().GetQueryResult(queryString)
    if err != nil {
        return nil, fmt.Errorf("failed to get asset %s: %v", assetID, err)
    }
    if assetBytes == nil {
        return nil, fmt.Errorf("Tran with id %s and create date %v does not exist",assetID, createDate)
    }
    var asset Asset
    err = json.Unmarshal(assetBytes, &asset)
    if err != nil {return nil, err}
    return &asset, nil
}
