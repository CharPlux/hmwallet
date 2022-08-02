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

	// eth token
	contract string
}

func newETH(key *Key) Wallet {
	return &eth{
		name:   "Ethereum",
		symbol: "ETH",
		key:    key,
	}
}

func (c *eth) GetType() uint32 {
	return c.key.Opt.CoinType
}

func (c *eth) GetName() string {
	return c.name
}

func (c *eth) GetS