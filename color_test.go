package color

import (
	"fmt"
	"testing"
)

// Clearly not a real test - just a visual reference
func TestOutput(t *testing.T) {
	Init()
	fmt.Println(Red + "Hello" + Reset)
	fmt.Println(Green + "Hello" + Reset)
	fmt.Println(Yellow + "Hello" + Reset)
	fmt.Println(Blue + "Hello" + Reset)
	fmt.Println(Purple + "Hello" + Reset)
	fmt.Println(Cyan + "Hello" + Reset)
	fmt.Println(Gray + "Hello" + Reset)
	fmt.Println(White + "Hello" + Reset)
	fmt.Println(Red + "H" + Green + "el" + Purple + "lo" + Reset)
	fmt.Println(Ize(Red, "test"))
	fmt.Println(Ize(Bold+Red, "bolded"))
}

func TestColorize(t *testing.T) {
	Init()
	if Ize(Red, "red") != "\033[31mred\033[0m" {
		t.Fail()
	}
	if Ize(Green, "green") != "\033[32mgreen\033[0m" {
		t.Fail()
	}
	if Ize(Blue, "blue") != "\033[34mblue\033[0m" {
		t.Fail()
	}
}
