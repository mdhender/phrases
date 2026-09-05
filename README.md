# phrases

> [!IMPORTANT]
> **This module has moved to
> [`github.com/maloquacious/phrases`](https://github.com/maloquacious/phrases).**
>
> This repository is archived and will not receive further updates.

Phrases is a passphrase generator.

# Migrating

```sh
go get github.com/maloquacious/phrases@latest
```

Two changes to the API:

- `Generate` now draws from `crypto/rand`. This module used `math/rand/v2`,
  which documents its own output as potentially predictable and so is not a
  sound source for a value whose only job is to be a secret.
- `Generate` returns `(string, error)`. It reports `ErrTooFewWords` when asked
  for fewer than one word, and otherwise errors only if the system entropy
  source fails.

Before:

```go
import psg "github.com/mdhender/phrases"

fmt.Println(psg.Generate(5, "."))
```

After:

```go
import psg "github.com/maloquacious/phrases"

passphrase, err := psg.Generate(7, ".")
if err != nil {
	log.Fatal(err)
}
fmt.Println(passphrase)
```

The example above also moves from five words to seven. At about 10.3 bits of
entropy per word, five words give roughly 51.7 bits; 64 bits needs at least
`64 / 10.3 = 6.2` words.

# Words

List is derived from <https://www.eff.org/files/2016/09/08/eff_short_wordlist_1.txt>.
The list generates about 10.3 bits of entropy per word.
To get 64 bits of entropy, we need at least 64 / 10.3 = 6.2 words.
