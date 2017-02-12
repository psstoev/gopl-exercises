//!+

package weightconv

// KToP converts a Metric weight to Imperial.
func KToP(k Kilogram) Pound { return Pound(k / KilogramsInPound) }

// PToK converts an Imperial weight to Metric.
func PToK(p Pound) Kilogram { return Kilogram(p) * KilogramsInPound }

//!-
