package main

import "fmt"

func main() {
	// simple declaration
	var number int
	number = 25
	fmt.Println(number)
	// declaration automatic typing
	name := "Mike"
	fmt.Println("Variable nombre: " + name)
	// asigned a lot of variables at the same time
	name, number = "Jhon", 50
	fmt.Println(name, number)
	// change variables values
	nameToSwitch := "David"
	fmt.Println("nameBeforeChange: "+name, "nameToChangeBeforeChange: "+nameToSwitch)
	name, nameToSwitch = nameToSwitch, name
	fmt.Println("nameAfterChange: "+name, "nameToCHangeAfterChange: "+nameToSwitch)
}
