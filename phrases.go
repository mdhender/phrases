////////////////////////////////////////////////////////////////////////////////
// phrases - a passphrase generator
// Copyright (c) 2022 Michael D. Henderson
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
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
