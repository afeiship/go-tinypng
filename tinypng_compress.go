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

type CompressResult struct {
	Input struct {
		Size int    `json:"size"`
		Type string `json:"type"`
	} `json:"input"`
	Output struct {
		Size   int     `json:"size"`
		Type   string  `json:"type"`
		Width  int     `json:"width"`
		Height int     `json:"height"`
		Ratio  float64 `json:"ratio"`
		Url    string  `json:"url"`
	} `json:"output"`
}

func (c *Client) CompressFromFile(filepath string) (string, error) {

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

func (c *Client) CompressFromURL(url string) (string, error) {
	apiUrl := fmt.Sprintf("%s/shrink", BaseUrl)
	auth := fmt.Sprintf("api:%s", c.ApiKey)
	authHeader := base64.StdEncoding.EncodeToString([]byte(auth))

	return fetch.Post(apiUrl, &fetch.Config{
		DataType: fetch.JSON,
		Headers: map[string]string{
			"Authorization": "Basic " + authHeader,
		},
		Body: map[string]any{
			"source": map[string]string{
				"url": url,
			},
		},
	})
}
