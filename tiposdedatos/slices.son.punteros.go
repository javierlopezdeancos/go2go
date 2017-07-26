package main

import (
	"fmt"
	 "../helpers"
)

func main() {

	title := `*                  Los Slices son  punteros                    *`

	headers := helpers.GetHeaders()

	fmt.Println(headers.Break)
	fmt.Println(headers.Title)
	fmt.Println(headers.TitleSide)
	fmt.Println(headers.TitleSide)
	fmt.Println(title)
	fmt.Println(headers.TitleSide)
	fmt.Println(headers.TitleSide)
	fmt.Println(headers.Title)

	jedis := []string{"Anakin", "Darth Vader", "Luke", "Yoda"}
	darkSide := jedis[0:2]
	lightSide := jedis[1:4]

	fmt.Println(headers.Given)
	fmt.Println("Dado este array jedis := []string{'Anakin', 'Darth Vader', 'Luke', 'Yoda'},  jedi = ", jedis)

	fmt.Println(headers.Break)

	fmt.Println("Estos serian el slice del lado oscuro darkSide := jedis[0:2], darkSide = ", darkSide)
	fmt.Println("Estos serían el slice del lado luminoso lightSide := jedis[1:4], lightSide = ", lightSide)

	darkSide[1] = "Darth Maul"

	fmt.Println(headers.When)
	fmt.Println("Cambiando un elemento de un slice, en realidad estaremos cambiando el valor en la posición del")
	fmt.Println("array al que está apuntado")
	fmt.Println("\nSi cambiamos Darth Vader por Darth Maul en el slice del Lado Oscuro...")
	fmt.Println(headers.Then)
	fmt.Println("Ahora el array Jedis es: ", jedis)
}
