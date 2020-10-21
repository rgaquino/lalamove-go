package lalamove

import (
	"context"
	"fmt"
)

// GetQuotation ...
func (c *Client) GetQuotation(ctx context.Context, req *GetQuotationRequest) (resp *GetQuotationResponse, err error) {
	if err := c.post(ctx, "/v2/quotations", req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// PlaceOrder ...
func (c *Client) PlaceOrder(ctx context.Context, req *PlaceOrderRequest) (resp *PlaceOrderResponse, err error) {
	if err := c.post(ctx, "/v2/orders", req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// OrderDetails ...
func (c *Client) OrderDetails(ctx context.Context, req *OrderDetailsRequest) (resp *OrderDetailsResponse, err error) {
	if err := c.get(ctx, fmt.Sprintf("/v2/orders/%s", req.OrderID), nil, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// CancelOrder cancels the order based on teh Lalamove cancellation policy.
// Attempts to cancel an order that does not comply with the cancellation policy
// will get ERR_CANCELLATION_FORBIDDEN as response.
func (c *Client) CancelOrder(ctx context.Context, req *CancelOrderRequest) (resp *CancelOrderResponse, err error) {
	if err := c.get(ctx, fmt.Sprintf("/v2/orders/%s/cancel", req.OrderID), nil, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// DriverDetails retrieves the driver's information.
func (c *Client) DriverDetails(ctx context.Context, req *DriverDetailsRequest) (resp *DriverDetailsResponse, err error) {
	if err := c.get(ctx, fmt.Sprintf("/v2/orders/%s/drivers/%s", req.OrderID, req.DriverID), nil, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// GetDriverLocation retrieves driver's latest location in latitude and longitude.
// This information is available starting 1 hour prior to datetime specified in
// scheduleAt datetime and remain accessible until the order is completed.
// Attempts made outside of this time window will get 403 Forbidden response.
func (c *Client) GetDriverLocation(ctx context.Context, req *DriverLocationRequest) (resp *DriverDetailsResponse, err error) {
	if err := c.get(ctx, fmt.Sprintf("/v2/orders/%s/drivers/%s/location", req.OrderID, req.DriverID), nil, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
