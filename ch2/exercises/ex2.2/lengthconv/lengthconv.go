//!+

// Package lengthconv performs Metric length to Imperial length conversions.
package lengthconv

import "fmt"

type Meter float64
type Yard float64

const (
	MetersInYard Meter = 0.9144
)

func (m Meter) String() string { return fmt.Sprintf("%g m", m) }
func (y Yard) String() string {
	if y == 1 {
		return fmt.Sprintf("%g yd", y)
	} else {
		return fmt.Sprintf("%g yds", y)
	}
}

//!-
