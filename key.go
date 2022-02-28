
package hdwallet

import (
	"crypto/ecdsa"
	"encoding/hex"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/cpacia/bchutil"
)

// Key struct
type Key struct {
	Opt      *Options
	Extended *hdkeychain.ExtendedKey

	// for btc
	Private *btcec.PrivateKey
	Public  *btcec.PublicKey

	// for eth
	PrivateECDSA *ecdsa.PrivateKey
	PublicECDSA  *ecdsa.PublicKey
}

// NewKey creates a master key
// params: [Mnemonic], [Password], [Language], [Seed]
func NewKey(opts ...Option) (*Key, error) {
	var (
		err error
		o   = newOptions(opts...)
	)

	if len(o.Seed) <= 0 {
		o.Seed, err = NewSeed(o.Mnemonic, o.Password, o.Language)
	}

	if err != nil {
		return nil, err
	}

	extended, err := hdkeychain.NewMaster(o.Seed, o.Params)
	if err != nil {
		return nil, err
	}

	key := &Key{
		Opt:      o,
		Extended: extended,
	}

	err = key.init()
	if err != nil {
		return nil, err
	}

	return key, nil
}
