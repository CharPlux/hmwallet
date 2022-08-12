package hdwallet

func init() {
	coins[IOST] = newIOST
}

type iost struct {
	*eth
}

func newIO