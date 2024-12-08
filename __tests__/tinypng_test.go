package tinypng

import (
	"github.com/afeiship/go-tinypng"
	"os"
	"testing"
)

var apiKey = os.Getenv("TINYPNG_API_KEY")
var client = tinypng.New(apiKey)

func TestCompress(f *testing.T) {
	res, err := client.Compress("./test.png")
	if err != nil {
		f.Fatal(err)
	}

	f.Log(res)
}
