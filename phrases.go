////////////////////////////////////////////////////////////////////////////////
// phrases - a passphrase generator
// Copyright (c) 2022 Michael D. Henderson
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
////////////////////////////////////////////////////////////////////////////////

// Package phrases implements a passphrase generator inspired by https://xkcd.com/936/.
package phrases

import (
	crand "crypto/rand"
	"encoding/binary"
	"encoding/json"
	"github.com/pkg/errors"
	"math/rand"
	"os"
)

// NewGenerator returns an initialized passphrase generator.
func NewGenerator(wordlist, separators string) (*Generator, error) {
	var seed [8]byte
	if _, err := crand.Read(seed[:]); err != nil {
		return nil, err
	}
	g := &Generator{
		rnd: rand.New(rand.NewSource(int64(binary.LittleEndian.Uint64(seed[:])))),
	}

	b, err := os.ReadFile(wordlist)
	if err != nil {
		return nil, err
	} else if err := json.Unmarshal(b, &g.words); err != nil {
		return nil, err
	} else if len(g.words) < 2048 {
		return nil, errors.New("too few words")
	}

	for _, r := range separators {
		g.separators = append(g.separators, string(r))
	}
	if len(g.separators) == 0 {
		switch g.rnd.Int() % 5 {
		case 0:
			g.separators = append(g.separators, " ")
		case 1:
			g.separators = append(g.separators, ".")
		case 2:
			g.separators = append(g.separators, "+")
		case 3:
			g.separators = append(g.separators, "-")
		default:
			g.separators = append(g.separators, " ", ".", "+", "-")
		}
	}

	return g, nil
}

// Generator implements a passphrase generator inspired by https://xkcd.com/936/
type Generator struct {
	rnd        *rand.Rand
	words      []string
	separators []string
}

// Generate returns a passphrase with the given number of words.
func (g *Generator) Generate(n int) (passphrase string) {
	for ; n > 0; n-- {
		if len(passphrase) != 0 {
			passphrase += g.separators[g.rnd.Int()%len(g.separators)]
		}
		passphrase += g.words[g.rnd.Int()%len(g.words)]
	}
	return passphrase
}
