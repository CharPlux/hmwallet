package hdwallet

func init() {
	coins[DASH] = newDASH
}

type dash struct {
	*b