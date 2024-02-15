package httputil

import (
	"encoding/json"
	"io"
	"net/http"
)

// Get from reqURL
func Get(reqURL string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, err
	}

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

// GetWithJSONUnmarshal get and json.Unmarshal
func GetWithJSONUnmarshal(reqURL string, v any) error {
	raw, err := Get(reqURL)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(raw, v); err != nil {
		return err
	}

	return nil
}
