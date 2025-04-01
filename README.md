# しずかなインターネットのAPIクライアント

## usage
```go
package main

import (
	"fmt"
	"os"

	qi "github.com/yagihash/quietinternet"
)

func main() {
	c := qi.New("token")
	res, err := c.GetPost("slug")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println(res)
	}
}
```
