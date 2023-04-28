package util

import (
	"io"
	"net/http"
)

func GetReq(url string) []byte {
	client := &http.Client{}

	resp, _ := client.Get(url)

	body, _ := io.ReadAll(resp.Body)
	return body
}
