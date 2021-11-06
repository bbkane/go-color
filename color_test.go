package color_test

import (
	"fmt"
	"testing"

	"github.com/bbkane/go-color"
)

// Clearly not a real test - just a visual reference
// func TestOutput(t *testing.T) {
// 	c.Enable()
// 	fmt.Println(c.Red + "Hello" + c.Reset)
// 	fmt.Println(c.Green + "Hello" + c.Reset)
// 	fmt.Println(c.Yellow + "Hello" + c.Reset)
// 	fmt.Println(c.Blue + "Hello" + c.Reset)
// 	fmt.Println(c.Purple + "Hello" + c.Reset)
// 	fmt.Println(c.Cyan + "Hello" + c.Reset)
// 	fmt.Println(c.Gray + "Hello" + c.Reset)
// 	fmt.Println(c.White + "Hello" + c.Reset)
// 	fmt.Println(c.Red + "H" + c.Green + "el" + c.Purple + "lo" + c.Reset)
// 	fmt.Println(c.Add(c.Red, "test"))
// 	fmt.Println(c.Add(c.Bold+c.Red, "bolded"))
// 	fmt.Println(c.Add(c.Underline+c.Green, "green underline"))
// 	fmt.Println(c.Add(c.Inverse+c.Purple, "inverse purple"))
// }

// func TestColorize(t *testing.T) {
// 	c.Enable()
// 	if c.Add(c.Red, "red") != "\033[31mred\033[0m" {
// 		t.Fail()
// 	}
// 	if c.Add(c.Green, "green") != "\033[32mgreen\033[0m" {
// 		t.Fail()
// 	}
// 	if c.Add(c.Blue, "blue") != "\033[34mblue\033[0m" {
// 		t.Fail()
// 	}
// }

func TestNewColors(t *testing.T) {
	err := color.Enable()
	if err != nil {
		t.Fatalf("errorr in Enable: %v\n", err)
	}
	fmt.Println(color.Add(color.Bold+color.Underline+color.BrightForegroundCyan, "BackgroundCyan"))
	fmt.Println(color.Add(color.ForegroundRed, "ForegroundRed"))

}
