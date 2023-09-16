package client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

const requestTimeout = time.Second * 6

type Client struct {
	hc *http.Client
}

func New() *Client {
	return &Client{
		hc: &http.Client{
			Timeout: requestTimeout,
		},
	}
}

func (c *Client) SendPost(url string, data interface{}, secret string) error {
	jsonBody, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	if secret != "" {
		req.Header.Set(fiber.HeaderAuthorization, secret)
	}

	resp, err := c.hc.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return nil
}
