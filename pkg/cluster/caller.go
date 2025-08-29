package cluster

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func (c *cluster) doRequest(deadLine int, method, url string, body io.Reader, v interface{}) error {
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(deadLine) * time.Second)
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



func (c *cluster) Get(deadLine int, url string, v interface{}) error {
	return c.doRequest(deadLine, http.MethodGet, url, nil, v)
}

func (c *cluster) Post(deadLine int, url string, body interface{}, v interface{}) error {
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

func (c *cluster) Put(deadLine int, url string, body interface{}, v interface{}) error {
	var reader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return err
		}
		reader = bytes.NewReader(data)
	}
	return c.doRequest(deadLine, http.MethodPut, url, reader, v)
}

func (c *cluster) Delete(deadLine int, url string, v interface{}) error {
	return c.doRequest(deadLine, http.MethodDelete, url, nil, v)
}