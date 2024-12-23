package tinypng

import (
	"github.com/afeiship/go-tinypng"
	"os"
	"testing"
)

var apiKey = os.Getenv("TINYPNG_API_KEY")
var client = tinypng.New(apiKey)

func TestCompressFromFile(f *testing.T) {
	res, err := client.CompressFromFile("./test.png")
	if err != nil {
		f.Fatal(err)
	}

	f.Log(res)
}

func TestCompressFromURL(f *testing.T) {
	res, err := client.CompressFromURL("https://via.placeholder.com/150")
	if err != nil {
		f.Fatal(err)
	}

	f.Log(res)
}
