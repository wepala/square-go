// This file was auto-generated by Fern from our API Definition.

package customers

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

// Returns details for a single customer.
func (c *Client) RetrieveCustomer(
	ctx context.Context,
	// The ID of the customer to retrieve.
	customerId string,
	opts ...option.RequestOption,
) (*api.RetrieveCustomerResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareup.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/customers/%v", customerId)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *api.RetrieveCustomerResponse
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

// Updates a customer profile. To change an attribute, specify the new value. To remove an attribute, specify the value as an empty string or empty object.
//
// As a best practice, you should include the `version` field in the request to enable [optimistic concurrency](https://developer.squareup.com/docs/working-with-apis/optimistic-concurrency) control. The value must be set to the current version of the customer profile.
//
// To update a customer profile that was created by merging existing profiles, you must use the ID of the newly created profile.
//
// You cannot use this endpoint to change cards on file. To make changes, use the [Cards API](api:Cards) or [Gift Cards API](api:GiftCards).
func (c *Client) UpdateCustomer(
	ctx context.Context,
	// The ID of the customer to update.
	customerId string,
	request *api.UpdateCustomerRequest,
	opts ...option.RequestOption,
) (*api.UpdateCustomerResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareup.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/customers/%v", customerId)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *api.UpdateCustomerResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPut,
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

// Deletes a customer profile from a business. This operation also unlinks any associated cards on file.
//
// As a best practice, you should include the `version` field in the request to enable [optimistic concurrency](https://developer.squareup.com/docs/working-with-apis/optimistic-concurrency) control. The value must be set to the current version of the customer profile.
//
// To delete a customer profile that was created by merging existing profiles, you must use the ID of the newly created profile.
func (c *Client) DeleteCustomer(
	ctx context.Context,
	// The ID of the customer to delete.
	customerId string,
	request *api.DeleteCustomerRequest,
	opts ...option.RequestOption,
) (*api.DeleteCustomerResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareup.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/customers/%v", customerId)

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *api.DeleteCustomerResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodDelete,
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