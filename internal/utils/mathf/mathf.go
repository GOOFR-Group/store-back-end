package mathf

// Clamp restricts the given value to the range defined by the given minimum and maximum values.
// Returns the minimum if the given value is less than the minimum, or the maximum if it's greater.
func Clamp(v, minimum, maximum float64) float64 {
	if v > maximum {
		return maximum
	}
	if v < minimum {
		return minimum
	}
	return v
}
