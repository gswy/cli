# cli
> A simple go-cli package that provides basic command line interaction and colorization.

## USE
1. First you need to import it into your own go project.
```shell
go get github.com/wanyun/cli
```
2. Then you can use it normally:

```go
package main

import (
	"fmt"
	"github.com/wanyun/cli"
)

func main() {
	cli.Success("Action success!")
	cli.Error("Action error!")
	cli.Warn("Action Warn")
	ask := cli.Ask("Place enter name")
	secret := cli.Secret("Place enter password")

	fmt.Println("")
	fmt.Printf("name is: %v", ask)
	fmt.Printf("password is %v", secret)
}
```