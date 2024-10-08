# phrases
Phrases is a passphrase generator

# Usage

```go
package main

import (
  "fmt"
  psg "github.com/mdhender/phrases/v2"
)

func main() {
    // the list of separators is optional
    fmt.Println(psg.Generate(5, "."))
}
```

# Words
List is derived from https://www.eff.org/files/2016/09/08/eff_short_wordlist_1.txt.
The list generates about 10.3 bits of entropy per word.
To get 64 bits of entropy, we need at least 64 / 10.3 = 6.2 words.
