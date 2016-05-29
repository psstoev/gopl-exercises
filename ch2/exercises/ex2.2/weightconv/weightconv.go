//!+

// Package weightconv performs Metric weight to Imperial weight conversions.
package weightconv

import "fmt"

type Kilogram float64
type Pound float64

const (
	KilogramsInPound Kilogram = 0.45359237
)

func (k Kilogram) String() string { return fmt.Sprintf("%g kg", k) }
func (p Pound) String() string {
	if p == 1 {
		return fmt.Sprintf("%g lb", p)
	} else {
		return fmt.Sprintf("%g lbs", p)
	}
}

//!-
