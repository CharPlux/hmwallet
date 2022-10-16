package hdwallet

func init() {
	coins[LTC] = newLTC
}

type ltc struct {
	*btc
}

func newLTC(key *Key) Wallet {
	key.Opt.Params = &LTCP