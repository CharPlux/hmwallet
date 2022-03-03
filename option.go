
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