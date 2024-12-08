package tinypng

import "log"

const BaseUrl = "https://api.tinypng.com"
const LogPrefix = "tinypng-client"

type Client struct {
	ApiKey string
}

func New(apiKey string) *Client {
	log.SetPrefix(LogPrefix)
	return &Client{ApiKey: apiKey}
}
