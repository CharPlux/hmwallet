package hdwallet

func init() {
	coins[DASH] = newDASH
}

type dash struct {
	*btc
}

func n