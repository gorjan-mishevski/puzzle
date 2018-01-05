package puzzle

// Currency - Normalized return data from BaseAgainsMultiPrice.
type Currency struct {
	base  string
	rates map[string]float64
}

// Base - Returns the base currency we are asking for rates. Ex: Base = "USD".
func (c *Currency) Base() string {
	return c.base
}

// Rates - Returns the rates from the base currency to the requested ones. Ex: Base "USD" to "BTH, LIT"
func (c *Currency) Rates() map[string]float64 {
	return c.rates
}
