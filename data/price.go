package data

type Price struct {
	Value        float64      `json:"value"`
	CurrencyCode CurrencyCode `json:"currency_code"`
}
