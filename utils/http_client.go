package utils

import (
	"io"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: 10 * time.Second,
}

func MakeRequest(url string) (data []byte, err error, statusCode int) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, err, resp.StatusCode
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err, resp.StatusCode
	}

	return body, nil, resp.StatusCode
}
