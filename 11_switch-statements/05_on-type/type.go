package main

import "fmt"

// switch on types
// -- normally we switch on value of variable
// -- go allows you to switch on type of variable

type contact struct {
	greeting string
	name     string
}

//SwitchOnType works with interfaces
//we'll learn more about interfaces later
func SwitchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case contact:
		fmt.Println("contact")
	default:
		fmt.Println(fmt.Println("unknown"))
	}
}

func main() {
	SwitchOnType(7)
	SwitchOnType("Mcleod")
	var t = contact{"Good to see you, ", "Tim"}
	SwitchOnType(t)
	SwitchOnType(t.greeting)
	SwitchOnType(t.name)
	var a interface{} = "string"
}
