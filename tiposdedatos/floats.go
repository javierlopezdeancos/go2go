package main

import "fmt"

func main() {

	// Datos tipo float
	var f32 float32     // 6 dígitos decimales de precision
	var f64 float64     // 15 dígitos decimales de precision
	var c64 complex64   // número complejo para float32
	var c128 complex128 // número complejo para float128

	//ejemplos
	fmt.Print("\n float32, float64, complex64, complex128 = ", f32, f64, c64, c128) //  0 0 (0+0i) (0+0i)

	f32 = .0156
	f64 = .0156

	fmt.Print("\n float32 * 8.565656 = ", f32*8.565656) // 0.13362423
	fmt.Print("\n float64 * 8.565656 = ", f64*8.565656) // 0.1336242336

	c64 = complex(5, 3)
	fmt.Print("\n complex64 = ", c64) // (5+3i)

	c128 = complex(5, 3)
	fmt.Print("\n complex128 = ", c128) // (5+3i)

	fmt.Print("\n complex64 * 85458.65 = ", c64*85458.65)   // (427293.25+256375.94i)
	fmt.Print("\n complex128 * 85458.65 = ", c128*85458.65) // (427293.25+256375.94999999998i)

	fmt.Print("\n Parte Real - de complex64 * 85458.65 = ", real(c64*85458.65))         // 427293.25
	fmt.Print("\n Parte imaginaria - de complex128 * 85458.65 = ", imag(c128*85458.65)) // 256375.94999999998

	fmt.Print("\n")

}
