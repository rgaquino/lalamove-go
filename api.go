package lalamove

import (
	"context"
	"fmt"
)

// GetQuotation requests a quotation.
func (c *Client) GetQuotation(ctx context.Context, city CityCode, req *GetQuotationRequest) (*GetQuotationResponse, error) {
	path := "/v2/quotations"
	resp := &GetQuotationResponse{}
	if err := c.post(ctx, city, path, req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// PlaceOrder creates a shipment order. The quotation received from GetQuotation and the same body used
// to get the quotation request should be merged in the request body.
func (c *Client) PlaceOrder(ctx context.Context, city CityCode, req *PlaceOrderRequest) (*PlaceOrderResponse, error) {
	path := "/v2/orders"
	resp := &PlaceOrderResponse{}
	if err := c.post(ctx, city, path, req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// OrderDetails retrieves the shipment order information.
func (c *Client) OrderDetails(ctx context.Context, city CityCode, orderID string) (*OrderDetailsResponse, error) {
	path := fmt.Sprintf("/v2/orders/%s", orderID)
	resp := &OrderDetailsResponse{}
	if err := c.get(ctx, city, path, nil, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// CancelOrder cancels the order based on the Lalamove cancellation policy. Attempts to cancel an order that
// does not comply with the cancellation policy will get ERR_CANCELLATION_FORBIDDEN as response.
func (c *Client) CancelOrder(ctx context.Context, city CityCode, orderID string) error {
	return c.put(ctx, city, fmt.Sprintf("/v2/orders/%s/cancel", orderID), nil, nil)
}

// DriverDetails retrieves the driver's information.
func (c *Client) DriverDetails(ctx context.Context, city CityCode, orderID, driverID string) (*DriverDetailsResponse, error) {
	path := fmt.Sprintf("/v2/orders/%s/drivers/%s", orderID, driverID)
	resp := &DriverDetailsResponse{}
	if err := c.get(ctx, city, path, nil, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// DriverLocation retrieves driver's latest location in latitude and longitude. This information is available starting
// 1 hour prior to datetime specified in scheduleAt datetime and remain accessible until the order is completed.
// Attempts made outside of this time window will get 403 Forbidden response.
func (c *Client) DriverLocation(ctx context.Context, city CityCode, orderID, driverID string) (*DriverLocationResponse, error) {
	path := fmt.Sprintf("/v2/orders/%s/drivers/%s/location", orderID, driverID)
	resp := &DriverLocationResponse{}
	if err := c.get(ctx, city, path, nil, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
