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
var _from common.Address
var _balance *big.Int
var _lastNonce uint64
var _chainId *big.Int

func Init(ctx context.Context, privateKey string, chainId string) {
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatalf("Init: %v", err)
	}
	_key = key
	_from = crypto.PubkeyToAddress(_key.PublicKey)
	log.Println("Using account", _from.Hex())
	_balance = new(big.Int)
	_chainId = new(big.Int)
	_chainId.SetString(chainId, 10)

	c := SchedulerChanFromContext(ctx)
	c <- callback(func(context.Context) error {
		defer func() {
			if r := recover(); r != nil {
				log.Println("Context error", r)
			}
		}()
		nc := CCFromContext(ctx)
		balance, err := nc.GetBalanceAt(ctx, _from, nil)
		if err != nil {
			return err
		}
		_balance.Set(balance)
		return nil
	})
	c <- callback(func(context.Context) error {
		defer func() {
			if r := recover(); r != nil {
				log.Println("Context error", r)
			}
		}()
		nc := CCFromContext(ctx)
		nonce, err := nc.PendingNonceAt(ctx, _from)
		if err != nil {
			return fmt.Errorf("failed to retrieve account nonce: %v", err)
		}
		_lastNonce = nonce
		return nil
	})
}

func SendWithValue(ctx context.Context, to common.Address, value *big.Int) (common.Hash, error) {
	nc := CCFromContext(ctx)
	auth := bind.NewKeyedTransactor(_key)

	if value == nil {
		value = new(big.Int)
	}
	gasPrice, err := nc.SuggestGasPrice(ctx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to suggest gas price: %v", err)
	}
	gasLimit := big.NewInt(21000)
	rawTx := types.NewTransaction(_lastNonce, to, value, gasLimit, gasPrice, nil)
	signedTx, err := auth.Signer(types.NewEIP155Signer(_chainId), _from, rawTx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("SendWithValue: %v", err)
	}
	err = nc.SendTransaction(ctx, signedTx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("SendTransaction: %v", err)
	}
	_lastNonce++
	return signedTx.Hash(), nil
}
