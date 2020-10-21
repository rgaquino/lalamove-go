package lalamove

import "time"

// ServiceType is the range of vehicles that Lalamove provides to cater to different needs at different cities.
type ServiceType string

// ServiceType enum
const (
	ServiceTypeCar          ServiceType = "CAR"
	ServiceTypeLalago       ServiceType = "LALAGO"
	ServiceTypeLalapro      ServiceType = "LALAPRO"
	ServiceTypeMinivan      ServiceType = "MINIVAN"
	ServiceTypeMotorcycle   ServiceType = "MOTORCYCLE"
	ServiceTypeMPV          ServiceType = "MPV"
	ServiceTypeTataAce7     ServiceType = "TATA7FT"
	ServiceTypeTataAce8     ServiceType = "TATA8FT"
	ServiceTypeThreeWheeler ServiceType = "THREE_WHEELER"
	ServiceTypeTruck175     ServiceType = "TRUCK175"
	ServiceTypeTruck330     ServiceType = "TRUCK330"
	ServiceTypeTruck550     ServiceType = "TRUCK550"
	ServiceTypeUV           ServiceType = "UV_FIORINO"
	ServiceTypeVan          ServiceType = "VAN"
	ServiceType4x4          ServiceType = "4X4"
)

// SpecialRequest ...
type SpecialRequest string

// SpecialRequest enum
const (
	SpecialRequest1HelperTier1             SpecialRequest = "1HELPER_TIER1"
	SpecialRequest1HelperTier2             SpecialRequest = "1HELPER_TIER3"
	SpecialRequest1HelperTier3             SpecialRequest = "1HELPER_TIER2"
	SpecialRequestAddAssistantTier1        SpecialRequest = "ADDITIONAL_ASSISTANT_TIER1"
	SpecialRequestAddAssistantTier2        SpecialRequest = "ADDITIONAL_ASSISTANT_TIER2"
	SpecialRequestAddAssistantTier3        SpecialRequest = "ADDITIONAL_ASSISTANT_TIER3"
	SpecialRequestInsulatedBag             SpecialRequest = "INSULATED_BAG"
	SpecialRequestUVVan                    SpecialRequest = "UV_VAN"
	SpecialRequestLalabag                  SpecialRequest = "LALABAG"
	SpecialRequestLalabagBig               SpecialRequest = "LALABAG_BIG"
	SpecialRequestDoor2Door                SpecialRequest = "DOOR2DOOR"
	SpecialRequestDoor2DoorDriver          SpecialRequest = "DOOR2DOOR_DRIVER"
	SpecialRequestDoor2DoorTruck330        SpecialRequest = "DOOR2DOOR_TRUCK330"
	SpecialRequestDoor2DoorTruck550        SpecialRequest = "DOOR2DOOR_TRUCK550"
	SpecialRequestDoor2Door1HelperTruck175 SpecialRequest = "DOOR2DOOR_1HELPER_TRUCK175"
	SpecialRequestDoor2Door1HelperTruck330 SpecialRequest = "DOOR2DOOR_1HELPER_TRUCK330"
	SpecialRequestDoor2Door1HelperTruck550 SpecialRequest = "DOOR2DOOR_1HELPER_TRUCK550"
	SpecialRequestDoor2Door2HelperTruck330 SpecialRequest = "DOOR2DOOR_2HELPER_TRUCK330"
	SpecialRequestDoor2Door2HelperTruck550 SpecialRequest = "DOOR2DOOR_2HELPER_TRUCK550"
	SpecialRequestCOD                      SpecialRequest = "COD"
	SpecialRequestPurchaseService          SpecialRequest = "PURCHASE_SERVICE"
	SpecialRequestPurchaseServiceTier2     SpecialRequest = "PURCHASE_SERVICE_TIER_2"
	SpecialRequestExtraHelper              SpecialRequest = "EXTRA_HELPER"
	SpecialRequestExtraHelperTruck175      SpecialRequest = "EXTRA_HELPER_TRUCK175"
	SpecialRequestRoundtripMotorcycle      SpecialRequest = "ROUNDTRIP_MOTORYCYCLE"
	SpecialRequestRoundtripTruck175        SpecialRequest = "ROUNDTRIP_TRUCK175"
	SpecialRequestRoundtripTruck330        SpecialRequest = "ROUNDTRIP_TRUCK330"
	SpecialRequestQueueingMotorcycle       SpecialRequest = "QUEUEING_MOTORCYCLE"
	SpecialRequestReturnTrip               SpecialRequest = "RETURNTRIP"
	SpecialRequestReturnTripLorry          SpecialRequest = "RETURNTRIP_LORRY"
	SpecialRequestLoadingService           SpecialRequest = "LOADING_SERVICE"
	SpecialRequestFoodService              SpecialRequest = "FOOD_SERVICE"
	SpecialRequestDriverCarries            SpecialRequest = "DRIVER_CARRIES"
	SpecialRequest1Assistant1To2Drops      SpecialRequest = "1ASSISTANT_1_MINUS_2DROPS"
	SpecialRequest1Assistant3To4Drops      SpecialRequest = "1ASSISTANT_3_MINUS_4DROPS"
	SpecialRequest1AssistantPlusDrops      SpecialRequest = "1ASSISTANT_5_PLUS_DROPS"
	SpecialRequestRestricted               SpecialRequest = "RESTRICTED"
	SpecialRequestMovingDriver             SpecialRequest = "MOVING_DRIVER"
	SpecialRequestMovingDriver1Helper      SpecialRequest = "MOVING_DRIVER_1HELPER"
	SpecialRequestMovingDriver2Helper      SpecialRequest = "MOVING_DRIVER_2HELPER"
	SpecialRequestMovingDriver1HelperVan   SpecialRequest = "MOVING_DRIVER_1HELPER_VAN"
	SpecialRequestMovingDriver2HelperVan   SpecialRequest = "MOVING_DRIVER_2HELPER_VAN"
	SpecialRequestTailgate                 SpecialRequest = "TAILGATE"
	SpecialRequestCovered                  SpecialRequest = "COVERED"
	SpecialRequestHelpBuy                  SpecialRequest = "HELP_BUY"
	SpecialRequestGroundFloor1Way          SpecialRequest = "GROUND_FLOOR_ONE_WAY"
	SpecialRequestGroundFloor1Way2         SpecialRequest = "GROUND_FLOOR_ONE_WAY_2"
	SpecialRequestUpstairDownstair1Way     SpecialRequest = "UPSTAIR_DOWNSTAIR_ONE_WAY"
	SpecialRequestUpstairDownstair1Way2    SpecialRequest = "UPSTAIR_DOWNSTAIR_ONE_WAY_2"
)

// OrderStatus ...
type OrderStatus string

// OrderStatus enum
const (
	// OrderStatusAssigningDriver - Trying to match shipment with a driver.
	OrderStatusAssigningDriver OrderStatus = "ASSIGNING_DRIVER"
	// OrderStatusOngoing - A driver has accepted the order.
	OrderStatusOngoing OrderStatus = "ON_GOING"
	// OrderStatusPickedUp - The driver has picked up the order.
	OrderStatusPickedUp OrderStatus = "PICKED_UP"
	// OrderStatusCompleted - The order has been delivered successfully and transaction has concluded.
	OrderStatusCompleted OrderStatus = "COMPLETED"
	// OrderStatusCanceled - User has canceled the order.
	OrderStatusCanceled OrderStatus = "CANCELED"
	// OrderStatusRejected - The order was matched and rejected twice by two drivers in a row.
	OrderStatusRejected OrderStatus = "REJECTED"
	// OrderStatusExpired - The order expired as no drivers accepted the order.
	OrderStatusExpired OrderStatus = "EXPIRED"
)

// Address ...
type Address struct {
	// DisplayString is the street address in plain text. Use remarks in DeliveryInfo for building, floor and flat.
	DisplayString string
	// Country is the country code of the address and must match with X-LLM-Country in the request headers.
	Country string
}

// Location ...
type Location struct {
	// Lat is the latitude
	Lat string `json:"lat"`
	// Lng is the longitude
	Lng string `json:"lng"`
}

// Waypoint ...
type Waypoint struct {
	Location  Location
	Addresses map[Locale]Address
}

// DeliveryInfo ...
type DeliveryInfo struct {
	// ToStop is the index of waypoint in stops this information associates with, has to be >= 1
	// since the first stop's Delivery Info is tided to requesterContact.
	ToStop int64 `json:"toStop"`
	// Contact is the contact person at the stop specified in ToStop.
	Contact Contact `json:"toContact"`
	// Remarks gives additional info about the delivery. eg. building, floor and flat.
	// Use newline \r\n for better readability.
	Remarks *string `json:"remarks,omitempty"`
}

// Contact ...
type Contact struct {
	// Name is the name of the contact person
	Name string `json:"name"`
	// Phone must be a valid phone number, validation varies for each country/region.
	Phone string `json:"phone"`
}

// GetQuotationRequest ...
type GetQuotationRequest struct {
	// ServiceType is the type of vehicle, availability varies for each country/region.
	ServiceType ServiceType `json:"serviceType"`
	// Stops is an array of Waypoints (minimum 2, maximum 10)
	Stops []Waypoint `json:"stops"`
	// Deliveries is an array of DeliveryInfos
	Deliveries []DeliveryInfo `json:"deliveries"`
	// RequesterContact is the contact person at pick up point aka stop[0].
	RequesterContact Contact `json:"requesterContact"`
	// ScheduleAt is the pick up time in UTC timezone and ISO 8601 format.
	// Omit this field if you are placing an immediate order.
	ScheduleAt *string `json:"scheduleAt,omitempty"`
	// SpecialRequests are special requests for the order, availability varies for each country/region.
	SpecialRequests *[]SpecialRequest `json:"specialRequests,omitempty"`
}

// GetQuotationResponse ...
type GetQuotationResponse struct {
	AmountCents int64  `json:"totalFee"`
	Currency    string `json:"totalFeeCurrency"`
}

// Price ...
type Price struct {
	AmountCents int64  `json:"amount"`
	Currency    string `json:"currency"`
}

// PlaceOrderRequest ...
type PlaceOrderRequest struct {
	QuotedPrice Price `json:"quotedTotalFee"`
	// SendSms is set to end delivery updates via SMS to the recipient,
	// or the recipient of the LAST STOP for multi-stop orders. Defaults to true.
	SendSms *bool `json:"sms"`
	GetQuotationRequest
}

// PlaceOrderResponse ...
type PlaceOrderResponse struct {
	// OrderID is the order id
	OrderID string `json:"orderRef"`
	// CustomerOrderID is a UUID order id (deprecated)
	CustomerOrderID string `json:"customerOrderId"`
}

// OrderDetailsRequest ...
type OrderDetailsRequest struct {
	OrderID string
}

// OrderDetailsResponse ...
type OrderDetailsResponse struct {
	Status   OrderStatus `json:"status"`
	Price    Price       `json:"price"`
	DriverID *string     `json:"driverId"`
}

// CancelOrderRequest ...
type CancelOrderRequest struct {
	OrderID string
}

// CancelOrderResponse ...
type CancelOrderResponse struct{}

// DriverDetailsRequest ...
type DriverDetailsRequest struct {
	OrderID  string
	DriverID string
}

// DriverDetailsResponse ...
type DriverDetailsResponse struct {
	Contact
	PlateNumber string `json:"plateNumber"`
	PhotoURL    string `json:"photo"`
}

// DriverLocationRequest ...
type DriverLocationRequest struct {
	OrderID  string
	DriverID string
}

// DriverLocationResponse ...
type DriverLocationResponse struct {
	Location  Location  `json:"location"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// ErrorResponse ...
type ErrorResponse struct {
	Error string `json:"message"`
}
