// This file was auto-generated by Fern from our API Definition.

package orders

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

// Search all orders for one or more locations. Orders include all sales,
// returns, and exchanges regardless of how or when they entered the Square
// ecosystem (such as Point of Sale, Invoices, and Connect APIs).
//
// `SearchOrders` requests need to specify which locations to search and define a
// [SearchOrdersQuery](entity:SearchOrdersQuery) object that controls
// how to sort or filter the results. Your `SearchOrdersQuery` can:
//
// Set filter criteria.
// Set the sort order.
// Determine whether to return results as complete `Order` objects or as
// [OrderEntry](entity:OrderEntry) objects.
//
// Note that details for orders processed with Square Point of Sale while in
// offline mode might not be transmitted to Square for up to 72 hours. Offline
// orders have a `created_at` value that reflects the time the order was created,
// not the time it was subsequently transmitted to Square.
func (c *Client) SearchOrders(
	ctx context.Context,
	request *api.SearchOrdersRequest,
	opts ...option.RequestOption,
) (*api.SearchOrdersResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareup.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/orders/search"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *api.SearchOrdersResponse
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

// Creates a new [order](entity:Order) that can include information about products for
// purchase and settings to apply to the purchase.
//
// To pay for a created order, see
// [Pay for Orders](https://developer.squareup.com/docs/orders-api/pay-for-orders).
//
// You can modify open orders using the [UpdateOrder](api-endpoint:Orders-UpdateOrder) endpoint.
func (c *Client) CreateOrder(
	ctx context.Context,
	request *api.CreateOrderRequest,
	opts ...option.RequestOption,
) (*api.CreateOrderResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareup.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/orders"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *api.CreateOrderResponse
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

// Retrieves an [Order](entity:Order) by ID.
func (c *Client) RetrieveOrder(
	ctx context.Context,
	// The ID of the order to retrieve.
	orderId string,
	opts ...option.RequestOption,
) (*api.RetrieveOrderResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareup.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/orders/%v", orderId)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *api.RetrieveOrderResponse
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

// Updates an open [order](entity:Order) by adding, replacing, or deleting
// fields. Orders with a `COMPLETED` or `CANCELED` state cannot be updated.
//
// An `UpdateOrder` request requires the following:
//
//   - The `order_id` in the endpoint path, identifying the order to update.
//   - The latest `version` of the order to update.
//   - The [sparse order](https://developer.squareup.com/docs/orders-api/manage-orders/update-orders#sparse-order-objects)
//     containing only the fields to update and the version to which the update is
//     being applied.
//   - If deleting fields, the [dot notation paths](https://developer.squareup.com/docs/orders-api/manage-orders#on-dot-notation)
//     identifying the fields to clear.
//
// To pay for an order, see
// [Pay for Orders](https://developer.squareup.com/docs/orders-api/pay-for-orders).
func (c *Client) UpdateOrder(
	ctx context.Context,
	// The ID of the order to update.
	orderId string,
	request *api.UpdateOrderRequest,
	opts ...option.RequestOption,
) (*api.UpdateOrderResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareup.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/orders/%v", orderId)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *api.UpdateOrderResponse
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
