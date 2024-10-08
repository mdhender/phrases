// Copyright (c) 2024 Michael D Henderson. All rights reserved.

// Package phrases implements a passphrase generator inspired by https://xkcd.com/936/.
package phrases

import (
	"math/rand/v2"
)

// Generate returns a passphrase with the given number of words.
func Generate(n int, separators ...string) (passphrase string) {
	if len(separators) == 0 {
		separators = []string{" ", ".", "+", "-"}
	}

	passphrase += words[rand.IntN(len(words))]
	for ; n > 0; n-- {
		passphrase += separators[rand.IntN(len(separators))]
		passphrase += words[rand.IntN(len(words))]
	}

	return passphrase
}
