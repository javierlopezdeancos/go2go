package main

import "fmt"

func main() {

	fmt.Println("\n Primer bucle FOR")

	// FOR
	// ejemplo uno
	i := 1
	for i <= 5 {
		fmt.Println(" ", i)
		i++
	}

	fmt.Println("\n Segundo bucle FOR")

	// mejor syntax
	for j := 5; j <= 10; j++ {
		fmt.Println(" ", j)
	}

	// DO WHILE WITH FOR AND IF
	fmt.Println("\n Tercer bucle DO WHILE (FOR + IF) ")

	k := 10
	for {
		fmt.Println(" ", k)
		if k >= 15 {
			break
		}
		k++
	}

	// IF
	fmt.Println("\n Imprimiendo el if")

	a := 6
	fmt.Println("\n \"a\" es igual a 6")

	if a = a + 1; a == 7 {
		fmt.Println("\n si \"a\" mas uno -> ahora \"a\" es igual a 7 - IF")
	}

	// ELSE
	fmt.Println("\n Imprimiendo el if")

	b := 5
	fmt.Println("\n \"b\" es igual a 5")

	if b = b + 1; b == 7 {
		fmt.Println("\n si \"b\" mas uno -> ahora \"b\" es igual a 7 - IF")
	} else {
		fmt.Println("\n si \"b\" mas uno -> ahora \"b\" es igual a 6 - ELSE")
	}

}
