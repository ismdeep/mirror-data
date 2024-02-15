package util

import (
	"context"
	"net/http"
	"strings"

	"github.com/ismdeep/log"
	"go.uber.org/zap"
)

// DivideURL divide url into <site, request-path, query>
func DivideURL(url string) (string, string, string) {
	idx1 := strings.Index(url, "://")
	idx2 := strings.Index(url[idx1+3:], "/")
	site := url[:idx1+3+idx2]
	requestPath := url[idx1+3+idx2:]

	if !strings.Contains(requestPath, "?") {
		return site, requestPath, ""
	}

	idx := strings.Index(requestPath, "?")

	return site, requestPath[:idx], requestPath[idx+1:]
}

// RealOriginURL get real origin url
func RealOriginURL(url string) (string, error) {
	req, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		log.WithContext(context.TODO()).Error("failed to create request", zap.Error(err))
		return "", err
	}

	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		log.WithContext(context.TODO()).Error("failed to get resp", zap.Error(err))
		return "", err
	}

	location := resp.Header.Get("Location")
	if location != "" {
		return location, nil
	}

	return url, nil
}
