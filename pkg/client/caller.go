package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/DKeshavarz/armis/internal/logger"
)



type client struct {
	client http.Client
	logger logger.Logger
}

func New()Client{
	return &client{
		client: http.Client{},
		logger: logger.New("Client"),
	}
}

func (c *client) doRequest(deadLine int, method, url string, body io.Reader, v interface{}) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(deadLine)*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// if resp.StatusCode < 200 || resp.StatusCode >= 300 {
	// 	return fmt.Errorf("bad status: %s", resp.Status)
	// }

	if v != nil {
		return json.NewDecoder(resp.Body).Decode(v)
	}
	return nil
}

func (c *client) Get(deadLine int, url string, v interface{}) error {
	return c.doRequest(deadLine, http.MethodGet, url, nil, v)
}

func (c *client) Post(deadLine int, url string, body interface{}, v interface{}) error {
	var reader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return err
		}
		reader = bytes.NewReader(data)
	}
	return c.doRequest(deadLine, http.MethodPost, url, reader, v)
}

func (c *client) Put(deadLine int, url string, body interface{}, v interface{}) error {
	var reader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return err
		}
		reader = bytes.NewReader(data)
	}
	err := c.doRequest(deadLine, http.MethodPut, url, reader, v)
	c.logger.Debug("after put", logger.Field{Key: "put", Value: v})
	return err
}

func (c *client) Delete(deadLine int, url string, v interface{}) error {
	return c.doRequest(deadLine, http.MethodDelete, url, nil, v)
}
