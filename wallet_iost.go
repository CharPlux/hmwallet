package hdwallet

func init() {
	coins[IOST] = newIOST
}

type iost struct {
	*eth
}

func newIOST(key *Key) Wallet {
	token :=