package bar

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
)

type Client struct {
	logger log.Logger
	client resty.Client
}

func NewBarClient(log log.Logger, client resty.Client) Client {
	return Client{
		logger: log,
		client: client,
	}
}

type Bar struct {
	Id string `json:"id"`
}

func (c *Client) GetBar(ctx context.Context, barId string) (*Bar, error) {
	bar := &Bar{}
	r := c.client.NewRequest()
	rr, err := r.
		SetPathParam("bar_id", barId).
		SetResult(bar).
		Get("/v1/bar/{bar_id}/test")

	if err != nil {
		return nil, err
	}

	if rr.IsError() {
		return nil, fmt.Errorf("we are having internal error in the response:%w", rr.Error())
	}

	return bar, nil
}
