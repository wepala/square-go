// This file was auto-generated by Fern from our API Definition.

package client

import (
	core "github.com/wepala/square-go/api/core"
	locations "github.com/wepala/square-go/api/locations"
	merchants "github.com/wepala/square-go/api/merchants"
	option "github.com/wepala/square-go/api/option"
	orders "github.com/wepala/square-go/api/orders"
	payouts "github.com/wepala/square-go/api/payouts"
	refunds "github.com/wepala/square-go/api/refunds"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header

	Orders    *orders.Client
	Payouts   *payouts.Client
	Refunds   *refunds.Client
	Merchants *merchants.Client
	Locations *locations.Client
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: core.NewCaller(
			&core.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header:    options.ToHeader(),
		Orders:    orders.NewClient(opts...),
		Payouts:   payouts.NewClient(opts...),
		Refunds:   refunds.NewClient(opts...),
		Merchants: merchants.NewClient(opts...),
		Locations: locations.NewClient(opts...),
	}
}
