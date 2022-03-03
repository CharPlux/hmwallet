
package hdwallet

import (
	"github.com/tyler-smith/go-bip39"
	"github.com/tyler-smith/go-bip39/wordlists"
)

func setLanguage(language string) {
	switch language {
	case ChineseSimplified:
		bip39.SetWordList(wordlists.ChineseSimplified)
	case ChineseTraditional:
		bip39.SetWordList(wordlists.ChineseTraditional)
	}
}

// NewMnemonic creates a random mnemonic
func NewMnemonic(length int, language string) (string, error) {
	setLanguage(language)

	if length < 12 {
		length = 12
	}
