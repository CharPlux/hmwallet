
package main

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/foxnut/go-hdwallet"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
)

var (
	mnemonic = "range sheriff try enroll deer over ten level bring display stamp recycle"
)

// utxo is an unspent transaction output of wallet
type utxo struct {
	key    *hdwallet.Key
	hash   string // hash of this utxo. example: f75f1b0810857dff6972767d745a57d0aeb77ae4845ef4f698a7bd50ec34a4b4 (BTC Testnet)
	index  uint32 // index of this wallet at this utxo
	amount int64  // satoshis amount of this utxo
}

func (u *utxo) ToTxIn() (*wire.TxIn, error) {
	hash, err := chainhash.NewHashFromStr(u.hash)
	if err != nil {
		return nil, err
	}
