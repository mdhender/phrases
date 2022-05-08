# phrases
Phrases is a passphrase generator

# Usage

```go
package main

import (
  "fmt"
  "github.com/mdhender/phrases"
)

func main() {
	psg, err := phrases.NewGenerator("words.json", " ")
	if err != nil {
		fmt.Println(err)
	} else {
    fmt.Println(psg.Generate(5))
  }
}
```
