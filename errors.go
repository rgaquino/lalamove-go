package lalamove

import "errors"

var (
	errCredentialsMissing = errors.New("API Key credentials missing")
	errBaseURLMissing     = errors.New("base URL missing")
)

var (
	// apiErrUnknownError - Default error
	apiErrUnknownError = errors.New("ERR_UNKNOWN")
	// apiErrInvalidCountry - Incorrect country
	apiErrInvalidCountry = errors.New("ERR_INVALID_COUNTRY")
	// apiErrInvalidParams - General validation error
	apiErrInvalidParams = errors.New("ERR_INVALID_PARAMS")
	// apiErrRequiredField - Missing required fields
	apiErrRequiredField = errors.New("ERR_REQUIRED_FIELD")
	// apiErrDeliveryMismatch - Stops and Deliveries mismatch
	apiErrDeliveryMismatch = errors.New("ERR_DELIVERY_MISMATCH")
	// apiErrInsufficientStops - Not enough stops, number of stops should be between 2 and 10
	apiErrInsufficientStops = errors.New("ERR_INSUFFICIENT_STOPS")
	// apiErrTooManyStops - Reached maximum stops, Number of stops should be between 2 and 10
	apiErrTooManyStops = errors.New("ERR_TOO_MANY_STOPS")
	// apiErrInvalidPaymentMethod - Invalid payment method
	apiErrInvalidPaymentMethod = errors.New("ERR_INVALID_PAYMENT_METHOD")
	// apiErrInvalidLocale - Invalid locale
	apiErrInvalidLocale = errors.New("ERR_INVALID_LOCALE")
	// apiErrInvalidPhoneNumber - Invalid phone number
	apiErrInvalidPhoneNumber = errors.New("ERR_INVALID_PHONE_NUMBER")
	// apiErrInvalidScheduleTime - scheduleAt datetime is in the past
	apiErrInvalidScheduleTime = errors.New("ERR_INVALID_SCHEDULE_TIME")
	// apiErrInvalidServiceType - No such service type, make sure to stick to service types that are available for the country/region
	apiErrInvalidServiceType = errors.New("ERR_INVALID_SERVICE_TYPE")
	// apiErrInvalidSpecialRequest - No such special request(s), make sure that special requests match with selected service type
	apiErrInvalidSpecialRequest = errors.New("ERR_INVALID_SPECIAL_REQUEST")
	// apiErrOutOfServiceArea - Out of service area
	apiErrOutOfServiceArea = errors.New("ERR_OUT_OF_SERVICE_AREA")
	// apiErrReverseGeocodeFailure - Fail to reverse from address to location, provide lat and lng
	apiErrReverseGeocodeFailure = errors.New("ERR_REVERSE_GEOCODE_FAILURE")
	// apiErrInsufficientCredit - You have insufficient credit, top up your wallet
	apiErrInsufficientCredit = errors.New("ERR_INSUFFICIENT_CREDIT")
	// apiErrInvalidCurrency - The currency you provided is not a valid currency
	apiErrInvalidCurrency = errors.New("ERR_INVALID_CURRENCY")
	// apiErrPriceMismatch - The amount or currency you provided in quotedTotalFee doesn't match quotation
	apiErrPriceMismatch = errors.New("ERR_PRICE_MISMATCH")
	// apiErrCancellationForbidden - Cancellation Forbidden
	apiErrCancellationForbidden = errors.New("ERR_CANCELLATION_FORBIDDEN")
	// apiErrTooManyRequests - Too many requests were made.
	apiErrTooManyRequests = errors.New("ERR_TOO_MANY_REQUESTS")
)

func wrapAPIError(errResp *ErrorResponse) error {
	switch errResp.Error {
	case "ERR_INVALID_COUNTRY":
		return apiErrInvalidCountry
	case "ERR_INVALID_PARAMS":
		return apiErrInvalidParams
	case "ERR_REQUIRED_FIELD":
		return apiErrRequiredField
	case "ERR_DELIVERY_MISMATCH":
		return apiErrDeliveryMismatch
	case "ERR_INSUFFICIENT_STOPS":
		return apiErrInsufficientStops
	case "ERR_TOO_MANY_STOPS":
		return apiErrTooManyStops
	case "ERR_INVALID_PAYMENT_METHOD":
		return apiErrInvalidPaymentMethod
	case "ERR_INVALID_LOCALE":
		return apiErrInvalidLocale
	case "ERR_INVALID_PHONE_NUMBER":
		return apiErrInvalidPhoneNumber
	case "ERR_INVALID_SCHEDULE_TIME":
		return apiErrInvalidScheduleTime
	case "ERR_INVALID_SERVICE_TYPE":
		return apiErrInvalidServiceType
	case "ERR_INVALID_SPECIAL_REQUEST":
		return apiErrInvalidSpecialRequest
	case "ERR_OUT_OF_SERVICE_AREA":
		return apiErrOutOfServiceArea
	case "ERR_REVERSE_GEOCODE_FAILURE":
		return apiErrReverseGeocodeFailure
	case "ERR_INSUFFICIENT_CREDIT":
		return apiErrInsufficientCredit
	case "ERR_INVALID_CURRENCY":
		return apiErrInvalidCurrency
	case "ERR_PRICE_MISMATCH":
		return apiErrPriceMismatch
	case "ERR_CANCELLATION_FORBIDDEN":
		return apiErrCancellationForbidden
	}
	return apiErrUnknownError
}
