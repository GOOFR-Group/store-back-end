package format

import (
	"fmt"
)

// Assemble returns a string by appending part to whole placing a space between them
func Assemble(whole, part string) string {
	return AssembleWith(whole, part, " ")
}

// AssembleWith returns a string by appending part to whole with separator between them
func AssembleWith(whole, part, separator string) string {
	if part == "" {
		return whole
	}
	if whole == "" {
		return part
	}
	return whole + separator + part
}

// AssembleFromKeyValue returns a string by appending key=value if value isn't empty, separating with a space
func AssembleFromKeyValue(whole, key, value string) string {
	if value == "" {
		return whole
	}
	return AssembleWith(whole, fmt.Sprintf("%s=%s", key, value), " ")
}
