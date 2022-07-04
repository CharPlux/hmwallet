
package hdwallet

func init() {
	coins[BTC] = newBTC
}

type btc struct {
	name   string
	symbol string