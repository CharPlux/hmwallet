
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
	return c.key.Opt.CoinType
}

func (c *etc) GetName() string {
	return c.name
}

func (c *etc) GetSymbol() string {
	return c.symbol