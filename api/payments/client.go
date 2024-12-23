// This file was auto-generated by Fern from our API Definition.

package payments

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

// Retrieves a list of payments taken by the account making the request.
//
// Results are eventually consistent, and new payments or changes to payments might take several
// seconds to appear.
//
// The maximum results per page is 100.
func (c *Client) ListPayments(
	ctx context.Context,
	request *api.ListPaymentsRequest,
	opts ...option.RequestOption,
) (*api.ListPaymentsResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareup.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/payments"

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *api.ListPaymentsResponse
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

// Creates a payment using the provided source. You can use this endpoint
// to charge a card (credit/debit card or
// Square gift card) or record a payment that the seller received outside of Square
// (cash payment from a buyer or a payment that an external entity
// processed on behalf of the seller).
//
// The endpoint creates a
// `Payment` object and returns it in the response.
func (c *Client) CreatePayment(
	ctx context.Context,
	request *api.CreatePaymentRequest,
	opts ...option.RequestOption,
) (*api.CreatePaymentResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareup.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/payments"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *api.CreatePaymentResponse
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

// Retrieves details for a specific payment.
func (c *Client) GetPayment(
	ctx context.Context,
	// A unique ID for the desired payment.
	paymentId string,
	opts ...option.RequestOption,
) (*api.GetPaymentResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareup.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/payments/%v", paymentId)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *api.GetPaymentResponse
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

// Updates a payment with the APPROVED status.
// You can update the `amount_money` and `tip_money` using this endpoint.
func (c *Client) UpdatePayment(
	ctx context.Context,
	// The ID of the payment to update.
	paymentId string,
	request *api.UpdatePaymentRequest,
	opts ...option.RequestOption,
) (*api.UpdatePaymentResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareup.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/payments/%v", paymentId)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *api.UpdatePaymentResponse
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
