
package hdwallet

func init() {
	coins[USDC] = newUSDC
}

type usdc struct {
	*eth
}

func newUSDC(key *Key) Wallet {
	token := newETH(key).(*eth)
	token.name = "USD Coin"
	token.symbol = "USDC"