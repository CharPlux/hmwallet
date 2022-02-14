
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