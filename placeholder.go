package gunship

import "fmt"

func Placeholder() {
	fmt.Println("Joe Mama Biden")
}

func ModifyGlobal(mod string) {
	Global = mod
}

var Global string = "Never gonna give you up!"

const VERSION = "0.0.1"
