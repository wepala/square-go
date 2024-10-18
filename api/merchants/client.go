// This file was auto-generated by Fern from our API Definition.

package merchants

import (
	context "context"
	api "github.com/wepala/square-go/api"
	core "github.com/wepala/square-go/api/core"
	option "github.com/wepala/square-go/api/option"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header
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
		header: options.ToHeader(),
	}
}

// Retrieves the `Merchant` object for the given `merchant_id`.
func (c *Client) RetrieveMerchant(
	ctx context.Context,
	// The ID of the merchant to retrieve. If the string "me" is supplied as the ID,
	// then retrieve the merchant that is currently accessible to this call.
	merchantId string,
	opts ...option.RequestOption,
) (*api.RetrieveMerchantResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareup.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/merchants/%v", merchantId)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *api.RetrieveMerchantResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}