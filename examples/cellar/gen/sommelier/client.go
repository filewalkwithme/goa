// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// sommelier client
//
// Command:
// $ goa gen goa.design/goa/examples/cellar/design -o
// $(GOPATH)/src/goa.design/goa/examples/cellar

package sommelier

import (
	"context"

	goa "goa.design/goa"
)

// Client is the "sommelier" service client.
type Client struct {
	PickEndpoint goa.Endpoint
}

// NewClient initializes a "sommelier" service client given the endpoints.
func NewClient(pick goa.Endpoint) *Client {
	return &Client{
		PickEndpoint: pick,
	}
}

// Pick calls the "pick" endpoint of the "sommelier" service.
// Pick can return the following error types:
//	- NoCriteria
//	- NoMatch
//	- error: generic transport error.
func (c *Client) Pick(ctx context.Context, p *Criteria) (res StoredBottleCollection, err error) {
	var ires interface{}
	ires, err = c.PickEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(StoredBottleCollection), nil
}