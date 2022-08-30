package hdwallet

func init() {
	coins[IOST] = newIOST
}

type iost struct {
	*eth
}

func newIOST(key *Key) Wallet {
	token := newETH(key).(*eth)
	token.name = "Internet of Services"
	token.symbol 