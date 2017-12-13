package internal

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	eth "github.com/ethereum/go-ethereum/params"

	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/Magicking/dimple/models"
	"github.com/Magicking/dimple/restapi/operations"
)

func ListHandler(ctx context.Context, params operations.ListParams) middleware.Responder {
	wrinkles, err := GetLimitedDB(ctx, 10)
	if err != nil {
		err_str := fmt.Sprintf("operations.ListHandlerFunc: %v", err)
		log.Println(err)
		return operations.NewListDefault(500).WithPayload(&models.Error{
			Message: &err_str})
	}

	payload := make([]*models.ListItem, len(wrinkles))
	for i, e := range wrinkles {
		payload[i] = &models.ListItem{Addr: e.To,
			Txid:   e.Txid,
			Amount: e.Amount,
		}
	}
	return operations.NewListOK().WithPayload(payload)
}
func SendHandler(ctx context.Context, params operations.SendParams) middleware.Responder {
	to := common.StringToAddress(params.To)
	amount := big.NewInt(eth.Ether)
	if _balance.Cmp(amount) < 0 {
		err_str := fmt.Sprintf("Insuficient funds: %v", _balance)
		return operations.NewSendDefault(500).WithPayload(&models.Error{
			Message: &err_str})
	}
	txid, err := SendWithValue(ctx, to, amount)
	if err != nil {
		err_str := fmt.Sprintf("operations.SendHandlerFunc: %v", err)
		log.Println(err)
		return operations.NewSendDefault(500).WithPayload(&models.Error{
			Message: &err_str})
	}

	go func() {
		wkl := &Wrinkle{To: to.Hex(), Txid: txid.Hex(), Amount: amount.String()}
		if err := InsertWrinkle(ctx, wkl); err != nil {
			log.Println("InsertWrinkle", err)
		}
	}()
	return operations.NewSendOK().WithPayload(txid.Hex())
}
