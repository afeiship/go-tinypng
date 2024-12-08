# go-tinypng
> Tinypng sdk for golang.

## installation
```sh
go get -u github.com/afeiship/go-tinypng
```

## usage
```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/afeiship/go-tinypng"
)

var apiKey = os.Getenv("TINYPNG_API_KEY")
var client = tinypng.New(apiKey)

func main() {
	// compress image from file
	result1, err := client.CompressFromFile("image.png")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result1)

	// compress image from url
	result2, err := client.CompressFromURL("https://example.com/image.png")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result2)
}
```

## resources
- https://tinypng.com/developers/reference