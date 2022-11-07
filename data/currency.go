package data

type CurrencyCode int

const (
	USD CurrencyCode = iota
	EUR
	GBP
	CAD
)

func (c CurrencyCode) String() string {
	switch c {
	case USD:
		return "USD"
	case EUR:
		return "EUR"
	case GBP:
		return "GBP"
	case CAD:
		return "CAD"
	default:
		return "Unknown"
	}
}

// handle JSON Marshalling
func (c CurrencyCode) MarshalJSON() ([]byte, error) {
	return []byte(`"` + c.String() + `"`), nil
}

// handle JSON Unmarshalling
func (c *CurrencyCode) UnmarshalJSON(b []byte) error {
	switch string(b) {
	case `"USD"`:
		*c = USD
	case `"EUR"`:
		*c = EUR
	case `"GBP"`:
		*c = GBP
	case `"CAD"`:
		*c = CAD
	default:
		*c = -1
	}
	return nil
}
