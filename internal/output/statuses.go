package output

import (
	"github.com/wayneashleyberry/truecolor/pkg/color"
)

// Output a success
func Success(message string) {
	color.Color(0, 255, 0).Println(message)
}