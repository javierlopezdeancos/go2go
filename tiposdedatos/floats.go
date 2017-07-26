package main

import (
	"fmt"
	"../helpers"
)

func main() {

	headers := helpers.GetHeaders()

	title := `*                       Los Floats                            *`

	fmt.Println(headers.Break)
	fmt.Println(headers.Title)
	fmt.Println(headers.TitleSide)
	fmt.Println(headers.TitleSide)
	fmt.Println(title)
	fmt.Println(headers.TitleSide)
	fmt.Println(headers.TitleSide)
	fmt.Println(headers.Title)

	var f32 float32
	var f64 float64
	var c64 complex64
	var c128 complex128

	fmt.Println(headers.Given)
	fmt.Print("Definiendo:  ")
	fmt.Print("var f32 float32, 6 dígitos decimales de precision")
	fmt.Print("var f64 float64, 15 dígitos decimales de precision")
	fmt.Print("var c64 complex64, número complejo para float32")
	fmt.Print("var c128 complex128, número complejo para float128")

	fmt.Println(headers.Then)
	fmt.Print("float32, float64, complex64, complex128 = ", f32, f64, c64, c128)

	fmt.Println(headers.Break)

	f32 = .0156
	f64 = .0156

	fmt.Println(headers.Given)
	fmt.Print("Definiendo:  ")
	fmt.Print("f32 = .0156")
	fmt.Print("f64 = .0156")

	fmt.Println(headers.Then)
	fmt.Print("f32*8.565656 = ", f32*8.565656)
	fmt.Print("f64*8.565656 = ", f64*8.565656)

	fmt.Println(headers.Break)

	c64 = complex(5, 3)

	fmt.Println(headers.Given)
	fmt.Print("Definiendo c64 = complex(5, 3)")

	fmt.Println(headers.Then)
	fmt.Print("c64 = ", c64)

	fmt.Println(headers.Break)

	c128 = complex(5, 3)

	fmt.Println(headers.Given)
	fmt.Print("Definiendo c128 = complex(5, 3)")

	fmt.Println(headers.Then)
	fmt.Print("c128 = ", c128)

	fmt.Println(headers.Then)
	fmt.Println("c64*85458.65 = ", c64*85458.65)
	fmt.Println("c128*85458.65 = ", c128*85458.65)

	fmt.Println("\nParte Real de c64*85458.65 se puede calcular con real(c64*85458.65) = ", real(c64*85458.65))
	fmt.Println("Parte Imaginaria de c128*85458.65 se puede calcular con imag(c128*85458.65) = ", imag(c128*85458.65))

}
