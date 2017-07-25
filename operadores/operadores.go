package main

import "fmt"

func main() {

	// Operaciones aritmeticas
	// Suma +
	// Resta -
	// Multiplicación *
	// Division /
	// Resto %

	a := 3
	b := 7

	fmt.Println("Operadores aritméticos ************************")
	fmt.Println("3 + 7 = ", a+b) // 10
	fmt.Println("3 - 7 = ", a-b) // -4
	fmt.Println("3 * 7 = ", a*b) // 21
	fmt.Println("7 / 3 = ", b/a) // 2
	fmt.Println("7 % 3 = ", b%a) //1

	// Operaciones de comparación
	// Igual que ==
	// Direfente a !=
	// Mayor que >
	// Menor que <
	// Mayor o igual que >=
	// Menor o igual que <=

	fmt.Println("Operadores de comparación *********************")
	fmt.Println("3 = 7 ", a == b)  // false
	fmt.Println("3 != 7 ", a != b) // true
	fmt.Println("3 > 7 ", a > b)   // false
	fmt.Println("7 < 3 ", b < a)   // false

	// Operaciones de asignacion
	// += | a += b | a = a + b
	// -= | a -= b | a =  a - b
	// *= | a *= b | a = a * b
	// /= | a /= b | a = a / b
	// %= | a %= b | a = a % b

	// Opradores lógicos
	// AND &&
	// OR ||
	// Negacion !

	fmt.Println("Operadores lógicos ****************************")
	fmt.Println("And && ----------------------------------------")
	fmt.Println("true and true ->", true && true)     // true
	fmt.Println("false and false ->", false && false) // false
	fmt.Println("true and false ->", true && false)   // false
	fmt.Println("false and true ->", false && true)   // false
	fmt.Println("Or || -----------------------------------------")
	fmt.Println("true or true ->", true || true)     // true
	fmt.Println("false or false ->", false || false) // false
	fmt.Println("true or false ->", true || false)   // true
	fmt.Println("false or true ->", false || true)   // true
	fmt.Println("Negacion ! ------------------------------------")
	fmt.Println("negacion - not true ->", !true) // false

	// Jerarquia de Operadores
	// ()
	// * / %
	// + -
	// == != <= >=
	// &&
	// ||

	fmt.Println("Jerarquia de operadores ***********************")
	fmt.Println("5 * (3 + 2) = ", 5*(3+2)) // 25

	// Asignadores de incremento y decrementp
	// ++ | a++ | a = a + 1
	// -- | a-- | a = a - 1
	// b + a++
	// b + ++a  / No exite en go

	fmt.Println("Asignadores de incremento y decremento")
	// es una sentencia, no una expresion, no se puede usar en expresiones
	a++
	fmt.Println("si a=3 y b=7 -> b + a++ = ", b+a) // 11

}
