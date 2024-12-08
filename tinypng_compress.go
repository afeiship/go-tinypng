package tinypng

import (
	"encoding/base64"
	"fmt"
	"github.com/afeiship/go-fetch"
	"log"
	"os"
)

//curl https://api.tinify.com/shrink \
//--user api:API_KEY \
//--data-binary @01.png

func (c *Client) Compress(filepath string) (string, error) {

	apiUrl := fmt.Sprintf("%s/shrink", BaseUrl)

	fileData, err := os.ReadFile(filepath)

	if err != nil {
		log.Fatal("Error reading file", err)
	}

	auth := fmt.Sprintf("api:%s", c.ApiKey)
	authHeader := base64.StdEncoding.EncodeToString([]byte(auth))

	return fetch.Post(apiUrl, &fetch.Config{
		DataType: fetch.OCTET_STREAM,
		Headers: map[string]string{
			"Authorization": "Basic " + authHeader,
		},
		Body: fileData,
	})
}
