package lalamove

import "errors"

var (
	errCredentialsMissing = errors.New("API Key credentials missing")
	errBaseURLMissing     = errors.New("base URL missing")
)

var (
	// errUnknownError - Default error
	errUnknownError = errors.New("ERR_UNKNOWN")
	// errInvalidCountry - Incorrect country
	errInvalidCountry = errors.New("ERR_INVALID_COUNTRY")
	// errInvalidParams - General validation error
	errInvalidParams = errors.New("ERR_INVALID_PARAMS")
	// errRequiredField - Missing required fields
	errRequiredField = errors.New("ERR_REQUIRED_FIELD")
	// errDeliveryMismatch - Stops and Deliveries mismatch
	errDeliveryMismatch = errors.New("ERR_DELIVERY_MISMATCH")
	// errInsufficientStops - Not enough stops, number of stops should be between 2 and 10
	errInsufficientStops = errors.New("ERR_INSUFFICIENT_STOPS")
	// errTooManyStops - Reached maximum stops, Number of stops should be between 2 and 10
	errTooManyStops = errors.New("ERR_TOO_MANY_STOPS")
	// errInvalidPaymentMethod - Invalid payment method
	errInvalidPaymentMethod = errors.New("ERR_INVALID_PAYMENT_METHOD")
	// errInvalidLocale - Invalid locale
	errInvalidLocale = errors.New("ERR_INVALID_LOCALE")
	// errInvalidPhoneNumber - Invalid phone number
	errInvalidPhoneNumber = errors.New("ERR_INVALID_PHONE_NUMBER")
	// errInvalidScheduleTime - scheduleAt datetime is in the past
	errInvalidScheduleTime = errors.New("ERR_INVALID_SCHEDULE_TIME")
	// errInvalidServiceType - No such service type, make sure to stick to service types that are available for the country/region
	errInvalidServiceType = errors.New("ERR_INVALID_SERVICE_TYPE")
	// errInvalidSpecialRequest - No such special request(s), make sure that special requests match with selected service type
	errInvalidSpecialRequest = errors.New("ERR_INVALID_SPECIAL_REQUEST")
	// errOutOfServiceArea - Out of service area
	errOutOfServiceArea = errors.New("ERR_OUT_OF_SERVICE_AREA")
	// errReverseGeocodeFailure - Fail to reverse from address to location, provide lat and lng
	errReverseGeocodeFailure = errors.New("ERR_REVERSE_GEOCODE_FAILURE")
	// errInsufficientCredit - You have insufficient credit, top up your wallet
	errInsufficientCredit = errors.New("ERR_INSUFFICIENT_CREDIT")
	// errInvalidCurrency - The currency you provided is not a valid currency
	errInvalidCurrency = errors.New("ERR_INVALID_CURRENCY")
	// errPriceMismatch - The amount or currency you provided in quotedTotalFee doesn't match quotation
	errPriceMismatch = errors.New("ERR_PRICE_MISMATCH")
	// errCancellationForbidden - Cancellation Forbidden
	errCancellationForbidden = errors.New("ERR_CANCELLATION_FORBIDDEN")
	// errTooManyRequests - Too many requests were made
	errTooManyRequests = errors.New("ERR_TOO_MANY_REQUESTS")
	// errUnauthorized - The provided authorization token is wrong
	errUnauthorized = errors.New("ERR_UNAUTHORIZED")
)

func wrapAPIError(errResp *ErrorResponse) error {
	switch errResp.Error {
	case "ERR_INVALID_COUNTRY":
		return errInvalidCountry
	case "ERR_INVALID_PARAMS":
		return errInvalidParams
	case "ERR_REQUIRED_FIELD":
		return errRequiredField
	case "ERR_DELIVERY_MISMATCH":
		return errDeliveryMismatch
	case "ERR_INSUFFICIENT_STOPS":
		return errInsufficientStops
	case "ERR_TOO_MANY_STOPS":
		return errTooManyStops
	case "ERR_INVALID_PAYMENT_METHOD":
		return errInvalidPaymentMethod
	case "ERR_INVALID_LOCALE":
		return errInvalidLocale
	case "ERR_INVALID_PHONE_NUMBER":
		return errInvalidPhoneNumber
	case "ERR_INVALID_SCHEDULE_TIME":
		return errInvalidScheduleTime
	case "ERR_INVALID_SERVICE_TYPE":
		return errInvalidServiceType
	case "ERR_INVALID_SPECIAL_REQUEST":
		return errInvalidSpecialRequest
	case "ERR_OUT_OF_SERVICE_AREA":
		return errOutOfServiceArea
	case "ERR_REVERSE_GEOCODE_FAILURE":
		return errReverseGeocodeFailure
	case "ERR_INSUFFICIENT_CREDIT":
		return errInsufficientCredit
	case "ERR_INVALID_CURRENCY":
		return errInvalidCurrency
	case "ERR_PRICE_MISMATCH":
		return errPriceMismatch
	case "ERR_CANCELLATION_FORBIDDEN":
		return errCancellationForbidden
	}
	return errUnknownError
}
