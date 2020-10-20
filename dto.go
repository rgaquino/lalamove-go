package lalamove

type ServiceType string

type SpecialRequest string

type Locale string

type CountryCode string

// Locale enum
const (
	BrasilEN      Locale = "en_BR"
	BrasilPT      Locale = "pt_BR"
	HongKongEN    Locale = "en_HK"
	HongKongZH    Locale = "zh_HK"
	IndiaEN       Locale = "en_IN"
	IndiaHI       Locale = "hi_IN"
	IndiaKN       Locale = "kn_IN"
	IndiaMR       Locale = "mr_IN"
	IndonesiaEN   Locale = "en_ID"
	IndonesiaID   Locale = "id_ID"
	MalaysiaEN    Locale = "en_MY"
	MalaysiaMS    Locale = "ms_MY"
	MexicoEN      Locale = "en_MX"
	MexicoMX      Locale = "es_MX"
	PhilippinesEN Locale = "en_PH"
	SingaporeEN   Locale = "en_SG"
	TaiwanZH      Locale = "zh_TW"
	ThailandEN    Locale = "en_TH"
	ThailandTH    Locale = "th_TH"
	VietnamEN     Locale = "en_VN"
	VietnamVI     Locale = "vi_VN"
)

// Address ...
type Address struct {
	// DisplayString is the street address in plain text. Use remarks in DeliveryInfo for building, floor and flat.
	DisplayString string
	// Country is the country code of the address and must match with X-LLM-Country in the request headers.
	Country CountryCode
}

// Location ...
type Location struct {
	// Lat is the latitude
	Lat string
	// Lng is the longitude
	Lng string
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
	ToStop int64
	// Contact is the contact person at the stop specified in ToStop.
	Contact Contact
	// Remarks gives additional info about the delivery. eg. building, floor and flat.
	// Use newline \r\n for better readability.
	Remarks *string
}

// Contact ...
type Contact struct {
	// Name is the name of the contact person
	Name string
	// Phone must be a valid phone number, validation varies for each country/region.
	Phone string
}

// GetQuotationRequest ...
type GetQuotationRequest struct {
	// ServiceType is the type of vehicle, availability varies for each country/region.
	ServiceType ServiceType
	// Stops is an array of Waypoints (minimum 2, maximum 10)
	Stops []Waypoint
	// Deliveries is an array of DeliveryInfos
	Deliveries []DeliveryInfo
	// RequesterContact is the contact person at pick up point aka stop[0].
	RequesterContact Contact
	// ScheduleAt is the pick up time in UTC timezone and ISO 8601 format.
	// Omit this field if you are placing an immediate order.
	ScheduleAt *string
	// SpecialRequests are special requests for the order, availability varies for each country/region.
	SpecialRequests *[]SpecialRequest
}

// GetQuotationResponse ...
type GetQuotationResponse struct {
}

// CreateOrderRequest ...
type CreateOrderRequest struct {
}

// CreateOrderResponse ...
type CreateOrderResponse struct {
}

// GetOrderDetailsRequest ...
type GetOrderDetailsRequest struct {
}

// GetOrderDetailsResponse ...
type GetOrderDetailsResponse struct {
}

// CancelOrderRequest ...
type CancelOrderRequest struct {
}

// CancelOrderResponse ...
type CancelOrderResponse struct {
}

// GetDriverDetailsRequest ...
type GetDriverDetailsRequest struct {
}

// GetDriverDetailsResponse ...
type GetDriverDetailsResponse struct {
}

// GetDriverLocationRequest ...
type GetDriverLocationRequest struct {
}

// GetDriverLocationResponse ...
type GetDriverLocationResponse struct {
}
