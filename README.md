# isupkg

## install
```bash
go get -u github.com/gari8/ispkg
```

## use
```go
package main

import (
	"github.com/gari8/ispkg"
)

func main() {
	p := isupkg.Profile{}
	_ = p.Run()
	defer p.Stop()
	// ...other codes
}
```