
package hdwallet

func init() {
	coins[BTC] = newBTC
}

type btc struct {
	name   string
	symbol string
	key    *Key
}

func newBTC(key *Key) Wallet {
	return &btc{
		name:   "Bitcoin",
		symbol: "BTC",
		key:    key,
	}
}

func (c *btc) GetType() uint32 {
	return c.key.Opt.CoinType
}
