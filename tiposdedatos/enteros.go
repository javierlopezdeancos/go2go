package main

import "fmt"
import "unsafe"

func main() {

	/********************************************************************************************/

	// Enteros con signo
	// var entero8 int8   //todos los enteros de 8 bits con signo (-128 to 127)
	// var entero16 int16 //todos los enteros de 16bits con signo (-32678 to 37677)
	// var entero32 int32 //todos los enteros de 32bits con signo (-2147483648 to 2147483647)

	// Enterps sin signo
	// var uentero8 uint8   //todos los enteros de 8bits  sin signo (0 to 255)
	// var uentero16 uint16 //todos los enteros de 16bits sin signo (0 to 65535)
	// var uentero32 uint32 //todos los enteros de 32 bits sin signo (0 to )

	// Alias
	// var enteroByte byte //alias para unit8
	// var enteroRune rune //alias para int32

	// Dependientes de la plataforma
	// var eneroUint uint        //32 a 64 bits
	// var enteroInt int         //32 a 64 bits
	// var enetroUintPtr uintptr // entero sin signo lo suficientemente grande para contener el valor de un puntero

	/********************************************************************************************/

	// Conversion entre tipos
	var entero32 int32
	var entero64 int64
	var enteroInt int

	entero32 = 25 // int32
	entero64 = 65 // uint 64

	enteroInt = 22 //32 o 64 bits, depende de la plataforma

	fmt.Println("suma") // 90
	fmt.Println(entero32 + int32(entero64))
	fmt.Println("tamaño de byten 64bits en bytes (64/8) ")
	fmt.Println(unsafe.Sizeof(enteroInt)) // 8
	fmt.Println("tamaño de int32 en bytes (32/8) ")
	fmt.Println(unsafe.Sizeof(entero32)) // 4
}
