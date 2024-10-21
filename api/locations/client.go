// This file was auto-generated by Fern from our API Definition.

package locations

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

// Provides details about all of the seller's [locations](https://developer.squareup.com/docs/locations-api),
// including those with an inactive status.
func (c *Client) ListLocations(
	ctx context.Context,
	opts ...option.RequestOption,
) (*api.ListLocationsResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareup.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/locations"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *api.ListLocationsResponse
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

// Creates a [location](https://developer.squareup.com/docs/locations-api).
// Creating new locations allows for separate configuration of receipt layouts, item prices,
// and sales reports. Developers can use locations to separate sales activity through applications
// that integrate with Square from sales activity elsewhere in a seller's account.
// Locations created programmatically with the Locations API last forever and
// are visible to the seller for their own management. Therefore, ensure that
// each location has a sensible and unique name.
func (c *Client) CreateLocation(
	ctx context.Context,
	request *api.CreateLocationRequest,
	opts ...option.RequestOption,
) (*api.CreateLocationResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareup.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/locations"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *api.CreateLocationResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
