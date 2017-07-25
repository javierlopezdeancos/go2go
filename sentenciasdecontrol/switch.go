package main

import "fmt"

func main() {

	var dia int

	error := `
    *************************************************************
    *                                                           *
    *                   1 Error encontrado                      *
    *                                                           *
    *************************************************************
    `
	header := `
    *************************************************************
    *                                                           *
    *                  Sentencias de control                    *
    *             Switch - bot días de la semana                *
    *                                                           *
    *************************************************************
    `
	footer := `
    *************************************************************
    *                                                           *
    *                 Resultado del ejercicio                   *
    *                                                           *
    *************************************************************
    `
	fmt.Println(header)
	fmt.Println("Introduce el número de la semana: ")
	fmt.Scanln(&dia)

	// Podemos hacerlo con ifs pero tambien con switch
	// if dia < 7 {
	// 	fmt.Println(footer)
	// 	if dia == 1 {
	// 		fmt.Println("Usted eligio el Lunes")
	// 	} else if dia == 2 {
	// 		fmt.Println("Usted eligio el Martes")
	// 	} else if dia == 3 {
	// 		fmt.Println("Usted eligio el Miercoles")
	// 	} else if dia == 4 {
	// 		fmt.Println("Usted eligio el Jueves")
	// 	} else if dia == 5 {
	// 		fmt.Println("Usted eligio el Viernes")
	// 	} else if dia == 6 {
	// 		fmt.Println("Usted eligio el Sabado")
	// 	} else if dia == 7 {
	// 		fmt.Println("Usted eligio el Domingo")
	// 	}
	// } else {
	// 	fmt.Println(error)
	// 	fmt.Println("Eligio un numero mayor que 7, Domingo")
	// }

	// en go no hace falta el break al final de cada case
	// el compilador lo entiende como caso de default
	switch dia {
	case 1:
		fmt.Println(footer)
		fmt.Println("Usted eligio el Lunes")
	case 2:
		fmt.Println(footer)
		fmt.Println("Usted eligio el Martes")
	case 3:
		fmt.Println(footer)
		fmt.Println("Usted eligio el Miercoles")
	case 4:
		fmt.Println(footer)
		fmt.Println("Usted eligio el Jueves")
	case 5:
		fmt.Println(footer)
		fmt.Println("Usted eligio el Viernes")
	case 6:
		fmt.Println(footer)
		fmt.Println("Usted eligio el Sabado")
	case 7:
		fmt.Println(footer)
		fmt.Println("Usted eligio el Domingo")
	default:
		fmt.Println(error)
		fmt.Println("Eligio un numero mayor que 7, Domingo")
	}

	// en go se puede usar el switch sin variable
	// realizando la comparación en el case
	switch {
	case dia <= 7:
		fmt.Println("La entrada introducida es correcta")
		// podemos no hacer break default y seguir comparando
		// con la palabra clave <fallthrough>
		// fallthrough
	case dia > 7:
		fmt.Println("La entrada introducida no es correcta")
	default:
		fmt.Println("Hemos tenido algún error")
	}

}
