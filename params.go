package hdwallet

import (
	"github.com/btcsuite/btcd/chaincfg"
)

// init net params
var (
	BTCParams        = chaincfg.MainNetParams
	BTCTestnetParams = chaincfg.T