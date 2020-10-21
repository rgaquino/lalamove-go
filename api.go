package lalamove

import (
	"context"
	"fmt"
)

// GetQuotation requests a quotation.
func (c *Client) GetQuotation(ctx context.Context, region UNLOCODE, req *GetQuotationRequest) (resp *GetQuotationResponse, err error) {
	if err := c.post(ctx, region, "/v2/quotations", req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// PlaceOrder creates a shipment order. The quotation received from GetQuotation and the same body used
// to get the quotation request should be merged in the request body.
func (c *Client) PlaceOrder(ctx context.Context, region UNLOCODE, req *PlaceOrderRequest) (resp *PlaceOrderResponse, err error) {
	if err := c.post(ctx, region, "/v2/orders", req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// OrderDetails retrieves the shipment order information.
func (c *Client) OrderDetails(ctx context.Context, region UNLOCODE, req *OrderDetailsRequest) (resp *OrderDetailsResponse, err error) {
	if err := c.get(ctx, region, fmt.Sprintf("/v2/orders/%s", req.OrderID), nil, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// CancelOrder cancels the order based on the Lalamove cancellation policy. Attempts to cancel an order that
// does not comply with the cancellation policy will get ERR_CANCELLATION_FORBIDDEN as response.
func (c *Client) CancelOrder(ctx context.Context, region UNLOCODE, req *CancelOrderRequest) (resp *CancelOrderResponse, err error) {
	if err := c.get(ctx, region, fmt.Sprintf("/v2/orders/%s/cancel", req.OrderID), nil, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// DriverDetails retrieves the driver's information.
func (c *Client) DriverDetails(ctx context.Context, region UNLOCODE, req *DriverDetailsRequest) (resp *DriverDetailsResponse, err error) {
	if err := c.get(ctx, region, fmt.Sprintf("/v2/orders/%s/drivers/%s", req.OrderID, req.DriverID), nil, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// DriverLocation retrieves driver's latest location in latitude and longitude. This information is available starting
// 1 hour prior to datetime specified in scheduleAt datetime and remain accessible until the order is completed.
// Attempts made outside of this time window will get 403 Forbidden response.
func (c *Client) DriverLocation(ctx context.Context, region UNLOCODE, req *DriverLocationRequest) (resp *DriverDetailsResponse, err error) {
	if err := c.get(ctx, region, fmt.Sprintf("/v2/orders/%s/drivers/%s/location", req.OrderID, req.DriverID), nil, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
