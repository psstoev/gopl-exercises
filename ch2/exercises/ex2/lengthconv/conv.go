//!+

package lengthconv

// MToY converts a Metric length to Imperial.
func MToY(m Meter) Yard { return Yard(m / MetersInYard) }

// YToM converts an Imperial length to Metric.
func YToM(y Yard) Meter { return Meter(y) * MetersInYard }

//!-
