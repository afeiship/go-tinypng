package tinypng

import (
	"github.com/afeiship/go-tinypng"
	"testing"
)

var client = tinypng.New("your_api_key")

func TestCompress(f *testing.T) {
	res, err := client.Compress("./test.png")
	if err != nil {
		f.Fatal(err)
	}

	f.Log(res)
}
