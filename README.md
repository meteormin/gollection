# gollection - go collection pkg

## install

```shell
go get meteormin/gollection
```

## usage

```go
package main

import (
	"github.com/meteormin/gollection"
	"log"
)

func main() {
	items := []int{1, 2, 3}
	collect := gollection.NewCollection(items)

	collect.For(func(v int, i int) {
		log.Print(v)
	})
	// result: 1 2 3

	collect.Chunk(100, func(v []int, i int) {
		log.Print(v) // chunked slice..
	})
}

```

- [collection_test.go](./collection_test.go)