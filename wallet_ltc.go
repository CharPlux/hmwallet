package hdwallet

func init() {
	coins[LTC] = newLTC
}

type ltc struct {
	*btc
}

func new