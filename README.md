# phrases
Phrases is a passphrase generator

# Usage

```go
package main

import (
  "fmt"
  psg "github.com/mdhender/phrases"
)

func main() {
    // the list of separators is optional
    fmt.Println(psg.Generate(5, "."))
}
```

# Words
Word listed copied from the Wikipedia article "Complete Shakespeare Wordlist."
