//go:build dev
// +build dev

package app

// init runs if dev flag is present
func init() {
	Dev = true
}
