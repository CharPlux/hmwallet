package hdwallet

func init() {
	coins[USDT] = newUSDT
}

type usdt struct {
	*btc
}

func ne