package hdwallet

func init() {
	coins[QTUM] = newQTUM
}

type qtum struct {
	