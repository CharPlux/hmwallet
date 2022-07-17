
package hdwallet

import (
	"github.com/ethereum/go-ethereum/crypto"
)

func init() {
	coins[ETC] = newETC
}

type etc struct {
	name   string
	symbol string
	key    *Key
}

func newETC(key *Key) Wallet {
	return &etc{
		name:   "Ethereum Classic",
		symbol: "ETC",
		key:    key,
	}
}

func (c *etc) GetType() uint32 {