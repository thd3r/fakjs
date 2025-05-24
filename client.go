package main

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

type Client struct {
	client *http.Client
}

func NewClient() *Client {
	var client Client

	client.client = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
			MaxConnsPerHost:     100,
			IdleConnTimeout:     time.Second,
			DisableKeepAlives:   true,
			TLSHandshakeTimeout: 15 * time.Second,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
			DialContext: (&net.Dialer{
				Timeout:   10 * time.Second,
				KeepAlive: time.Second,
			}).DialContext,
		},
		Timeout:       10 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
	}

	return &client
}

func (c *Client) Do(ctx context.Context, method, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", RandomAgents())
	req.Header.Set("Connection", "close")

	req.Close = true

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
