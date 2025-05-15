package bold

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const BASE_URL = "https://integrations.api.bold.co"

type Bold struct {
	ApiKey string
}

func NewClient(apiKey string) *Bold {
	return &Bold{
		ApiKey: apiKey,
	}
}

func (b *Bold) doRequest(method string, url string, body interface{}) ([]byte, error) {
	var buf io.Reader

	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewReader(jsonBody)
	}

	uri := fmt.Sprintf("%s/%s", BASE_URL, url)
	req, err := http.NewRequest(method, uri, buf)
	if err != nil {
		return nil, err
	}

	apiKey := fmt.Sprintf("x-api-key %s", b.ApiKey)

	req.Header.Set("Authorization", apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		log.Println("[Error] doRequest: ", resp.Status)
	}

	return io.ReadAll(resp.Body)
}
