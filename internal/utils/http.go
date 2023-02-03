package utils

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	connTimeout    = 5
	defaultTimeout = 10
)

var HTTPClient = &http.Client{
	Transport: transport(connTimeout),
	Timeout:   time.Second * time.Duration(defaultTimeout),
}

func transport(connTimeout int) *http.Transport {
	dialer := &net.Dialer{
		Timeout:   time.Second * time.Duration(connTimeout),
		KeepAlive: time.Second * 10,
	}
	return &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		MaxIdleConns:        100,
		MaxConnsPerHost:     200,
		MaxIdleConnsPerHost: 50,
		IdleConnTimeout:     600 * time.Second,
		DialContext:         dialer.DialContext,
	}
}

func HTTPGet(url string) ([]byte, error) {
	resp, err := HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status:%d, body:%s", resp.StatusCode, body)
	}

	return body, nil
}

func HTTPPost(url, contentType, reqBody string) ([]byte, error) {
	bodyReader := strings.NewReader(reqBody)
	resp, err := HTTPClient.Post(url, contentType, bodyReader)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status:%d, body:%s", resp.StatusCode, body)
	}

	return body, nil
}

func HTTPPostWithContext(ctx context.Context, url, contentType, reqBody string) ([]byte, error) {
	bodyReader := strings.NewReader(reqBody)
	req, err := http.NewRequest("POST", url, bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	resp, err := HTTPClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status:%d, body:%s", resp.StatusCode, body)
	}

	return body, nil
}
