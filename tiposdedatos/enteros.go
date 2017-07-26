package main

import (
	"fmt"
	"unsafe"
	"../helpers"
)

func main() {

	headers := helpers.GetHeaders()

	title := `*                       Los Enteros                            *`

	fmt.Println(headers.Break)
	fmt.Println(headers.Title)
	fmt.Println(headers.TitleSide)
	fmt.Println(headers.TitleSide)
	fmt.Println(title)
	fmt.Println(headers.TitleSide)
	fmt.Println(headers.TitleSide)
	fmt.Println(headers.Title)

	fmt.Println(headers.Given)
	fmt.Println("Definiendo:  ")

	fmt.Println("\nEnteros con signo")
	fmt.Println("var entero8 int8, todos los enteros de 8 bits con signo (-128 to 127)")
	fmt.Println("var entero16 int16, todos los enteros de 16bits con signo (-32678 to 37677)")
	fmt.Println("var entero32 int32, todos los enteros de 32bits con signo (-2147483648 to 2147483647)")

	fmt.Println("\nEnteros sin signo")
	fmt.Println("var uentero8 uint8, todos los enteros de 8bits  sin signo (0 to 255)")
	fmt.Println("var uentero16 uint16, todos los enteros de 16bits sin signo (0 to 65535)")
	fmt.Println("var uentero32 uint32, todos los enteros de 32 bits sin signo (0 to )")

	fmt.Println("\nAlias de enteros")
	fmt.Println("var enteroByte byte, alias para unit8")
	fmt.Println("var enteroRune rune, alias para int32")

	fmt.Println("\nDependientes de la plataforma")
	fmt.Println("var enteroUint uint, 32 a 64 bits")
	fmt.Println("var enteroInt int, 32 a 64 bits")
	fmt.Println("var enetroUintPtr uintptr, entero sin signo lo suficientemente grande para contener el valor de un puntero")

	var entero32 int32
	var entero64 int64
	var enteroInt int

	entero32 = 25
	entero64 = 65
	enteroInt = 22

	fmt.Println(headers.When)
	fmt.Println("entero32 = 25")
	fmt.Println("entero64 = 65")
	fmt.Println("enteroInt = 22")

	fmt.Println(headers.Then)
	fmt.Println("Suma")
	fmt.Println("entero32 + int32(entero64) = ", entero32 + int32(entero64))

	fmt.Println("\nTamaño de int 32 o 64 bits (depende de la plataforma) en bytes (64/8)")
	fmt.Println("unsafe.Sizeof(enteroInt) = ", unsafe.Sizeof(enteroInt))

	fmt.Println("\nTamaño de int32 en bytes (32/8) ")
	fmt.Println("unsafe.Sizeof(entero32) = ", unsafe.Sizeof(entero32))
}
