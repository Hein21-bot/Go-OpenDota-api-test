package clients

import (
	"encoding/json"
	"net/http"
	"time"
)

type OpenDotaClient struct {
	BaseURL string
	Client  *http.Client
}

func NewOpenDotaClient() *OpenDotaClient {
	return &OpenDotaClient{
		BaseURL: "https://api.opendota.com/api",
		Client: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

func (c *OpenDotaClient) Get(path string, target any) error {
	resp, err := c.Client.Get(c.BaseURL + path)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}
