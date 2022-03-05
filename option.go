
package hdwallet

import (
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/chaincfg"
)

// default options
var (
	DefaultParams = &BTCParams

	// master key options
	DefaultPassword = ""
	DefaultLanguage = ""

	// child key options
	DefaultPurpose      = ZeroQuote + 44
	DefaultCoinType     = BTC
	DefaultAccount      = ZeroQuote
	DefaultChange       = Zero
	DefaultAddressIndex = Zero
)

// Option of key
type Option func(*Options)

// Options of key
type Options struct {
	Params *chaincfg.Params

	// master key options
	Mnemonic string
	Password string
	Language string
	Seed     []byte

	// child key options
	Purpose      uint32
	CoinType     uint32
	Account      uint32
	Change       uint32
	AddressIndex uint32
}

func newOptions(opts ...Option) *Options {
	opt := &Options{
		Params:       DefaultParams,
		Password:     DefaultPassword,
		Language:     DefaultLanguage,
		Purpose:      DefaultPurpose,
		CoinType:     DefaultCoinType,
		Account:      DefaultAccount,
		Change:       DefaultChange,
		AddressIndex: DefaultAddressIndex,
	}

	for _, o := range opts {
		o(opt)
	}

	return opt
}

// GetPath return path in bip44 style
func (o *Options) GetPath() []uint32 {
	return []uint32{
		o.Purpose,
		o.CoinType,
		o.Account,
		o.Change,
		o.AddressIndex,
	}
}

// Params set to options
func Params(p *chaincfg.Params) Option {
	return func(o *Options) {
		o.Params = p
	}
}

// Mnemonic set to options
func Mnemonic(m string) Option {
	return func(o *Options) {
		o.Mnemonic = m
	}
}

// Password set to options
func Password(p string) Option {
	return func(o *Options) {
		o.Password = p
	}
}

// Language set to options
func Language(l string) Option {
	return func(o *Options) {
		o.Language = l
	}
}