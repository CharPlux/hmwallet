package hdwallet

import (
	"github.com/ethereum/go-ethereum/crypto"
)

func init() {
	coins[ETH] = newETH
}

type eth struct {
	name   string
	symbol string
	key    *Key

	// et