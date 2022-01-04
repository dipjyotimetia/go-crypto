package request

import (
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

type Connection interface {
	HTTPGet(url string, header map[string]string, param map[string]string) (*resty.Response, error)
	HTTPPost(url string, header map[string]string, param map[string]string, body interface{}) (*resty.Response, error)
}

type Client struct {
	client *resty.Client
}

func NewHTTPConn() Connection {
	client := resty.New()
	return &Client{
		client,
	}
}

func (c Client) HTTPGet(url string, header map[string]string, param map[string]string) (*resty.Response, error) {
	resp, err := c.client.R().
		SetHeaders(header).
		SetQueryParams(param).
		Get(url)
	if err != nil {
		log.Fatalf("request failed with error %s", err)
		return nil, err
	}
	return resp, nil
}

func (c Client) HTTPPost(url string, header map[string]string, param map[string]string, body interface{}) (*resty.Response, error) {
	resp, err := c.client.R().
		SetHeaders(header).
		SetQueryParams(param).
		SetBody(body).
		Post(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return resp, nil
}
