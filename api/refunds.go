// This file was auto-generated by Fern from our API Definition.

package api

type ListPaymentRefundsRequest struct {
	// The timestamp for the beginning of the requested reporting period, in RFC 3339 format.
	//
	// Default: The current time minus one year.
	BeginTime *string `json:"-" url:"begin_time,omitempty"`
	// The timestamp for the end of the requested reporting period, in RFC 3339 format.
	//
	// Default: The current time.
	EndTime *string `json:"-" url:"end_time,omitempty"`
	// The order in which results are listed:
	//
	// - `ASC` - Oldest to newest.
	// - `DESC` - Newest to oldest (default).
	SortOrder *string `json:"-" url:"sort_order,omitempty"`
	// A pagination cursor returned by a previous call to this endpoint.
	// Provide this cursor to retrieve the next set of results for the original query.
	//
	// For more information, see [Pagination](https://developer.squareup.com/docs/basics/api101/pagination).
	Cursor *string `json:"-" url:"cursor,omitempty"`
	// Limit results to the location supplied. By default, results are returned
	// for all locations associated with the seller.
	LocationId *string `json:"-" url:"location_id,omitempty"`
	// If provided, only refunds with the given status are returned.
	// For a list of refund status values, see [PaymentRefund](entity:PaymentRefund).
	//
	// Default: If omitted, refunds are returned regardless of their status.
	Status *string `json:"-" url:"status,omitempty"`
	// If provided, only returns refunds whose payments have the indicated source type.
	// Current values include `CARD`, `BANK_ACCOUNT`, `WALLET`, `CASH`, and `EXTERNAL`.
	// For information about these payment source types, see
	// [Take Payments](https://developer.squareup.com/docs/payments-api/take-payments).
	//
	// Default: If omitted, refunds are returned regardless of the source type.
	SourceType *string `json:"-" url:"source_type,omitempty"`
	// The maximum number of results to be returned in a single page.
	//
	// It is possible to receive fewer results than the specified limit on a given page.
	//
	// If the supplied value is greater than 100, no more than 100 results are returned.
	//
	// Default: 100
	Limit *int `json:"-" url:"limit,omitempty"`
}

type RefundPaymentRequest struct {
	//	A unique string that identifies this `RefundPayment` request. The key can be any valid string
	//
	// but must be unique for every `RefundPayment` request.
	//
	// Keys are limited to a max of 45 characters - however, the number of allowed characters might be
	// less than 45, if multi-byte characters are used.
	//
	// For more information, see [Idempotency](https://developer.squareup.com/docs/working-with-apis/idempotency).
	IdempotencyKey string `json:"idempotency_key" url:"-"`
	AmountMoney    *Money `json:"amount_money,omitempty" url:"-"`
	AppFeeMoney    *Money `json:"app_fee_money,omitempty" url:"-"`
	// The unique ID of the payment being refunded. Must be provided and non-empty.
	PaymentId *string `json:"payment_id,omitempty" url:"-"`
	// A description of the reason for the refund.
	Reason *string `json:"reason,omitempty" url:"-"`
	//	Used for optimistic concurrency. This opaque token identifies the current `Payment`
	//
	// version that the caller expects. If the server has a different version of the Payment,
	// the update fails and a response with a VERSION_MISMATCH error is returned.
	// If the versions match, or the field is not provided, the refund proceeds as normal.
	PaymentVersionToken *string `json:"payment_version_token,omitempty" url:"-"`
	// An optional [TeamMember](entity:TeamMember) ID to associate with this refund.
	TeamMemberId *string `json:"team_member_id,omitempty" url:"-"`
}
