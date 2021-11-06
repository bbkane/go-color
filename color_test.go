package color_test

import (
	"fmt"
	"testing"

	"github.com/bbkane/go-color"
)

// Run with go test . -v to see the output
func TestRed(t *testing.T) {
	color.Enable()
	fmt.Println(color.Add(color.ForegroundRed, "ForegroundRed"))
	fmt.Println(color.Add(color.BackgroundRed, "BackgroundRed"))
	fmt.Println(color.Add(color.BrightForegroundRed, "BrightForegroundRed"))
	fmt.Println(color.Add(color.BrightBackgroundRed, "BrightBacgroundRed"))
	fmt.Println(color.Add(color.Bold+color.BackgroundRed+color.ForegroundYellow, "mix"))

}
