package hdwallet

var coins = make(map[uint32]func(*Key) Wallet)

// Wallet interface
type Wallet interface {
	Get