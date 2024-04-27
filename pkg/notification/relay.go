package notification

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

func trimEndpoints(endpoints []string) []string {
	var results []string
	for _, endpoint := range endpoints {
		endpoint = strings.TrimSpace(endpoint)
		if endpoint != "" {
			results = append(results, endpoint)
		}
	}
	return results
}

func RelayMsg(relayEndpoints []string, token string, msg string) error {
	relayEndpoints = trimEndpoints(relayEndpoints)

	if len(relayEndpoints) <= 0 {
		return errors.New("relay endpoints empty")
	}

	if token == "" {
		return errors.New("token is empty")
	}

	if msg == "" {
		return errors.New("msg is empty")
	}

	msgUUID := uuid.NewString()

	var errLst []error

	var wg sync.WaitGroup
	for _, endpoint := range relayEndpoints {
		wg.Add(1)
		go func(requestURL string) {
			defer wg.Done()

			req, err := http.NewRequest(http.MethodPut, requestURL, bytes.NewBuffer([]byte(msg)))
			if err != nil {
				errLst = append(errLst, err)
				return
			}

			req.Header.Set("X-Relay-Auth", token)

			resp, err := (&http.Client{Timeout: time.Second * 5}).Do(req)
			if err != nil {
				errLst = append(errLst, err)
				return
			}

			if resp.StatusCode != http.StatusOK {
				errLst = append(errLst, fmt.Errorf("failed to send msg via %v, err: %v", requestURL, resp.Status))
				return
			}

			return
		}(fmt.Sprintf("%v/api/v1/msg/%v", endpoint, msgUUID))
	}
	wg.Wait()

	if len(errLst) >= len(relayEndpoints) {
		return errors.Join(errLst...)
	}

	return nil
}
