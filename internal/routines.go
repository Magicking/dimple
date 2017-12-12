package internal

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var _key *ecdsa.PrivateKey

func Init(ctx context.Context, privateKey string) {
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatalf("Init: %v", err)
	}
	_key = key
}

func SendWithValue(ctx context.Context, to common.Address, value *big.Int) (common.Hash, error) {
	nc := CCFromContext(ctx)
	auth := bind.NewKeyedTransactor(_key)
	from := crypto.PubkeyToAddress(_key.PublicKey)

	if value == nil {
		value = new(big.Int)
	}
	var nonce uint64
	nonce, err := nc.PendingNonceAt(ctx, from)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to retrieve account nonce: %v", err)
	} else {
		nonce = 0
	}
	gasPrice, err := nc.SuggestGasPrice(ctx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to suggest gas price: %v", err)
	}
	gasLimit := big.NewInt(21000)
	rawTx := types.NewTransaction(nonce, to, value, gasLimit, gasPrice, nil)
	signedTx, err := auth.Signer(types.HomesteadSigner{}, from, rawTx)
	if err != nil {
		log.Fatalf("SendWithValue: %v", err)
	}
	err = nc.SendTransaction(ctx, signedTx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("SendTransaction: %v", err)
	}
	return signedTx.Hash(), nil
}
