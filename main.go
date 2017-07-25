package main

import (
	"fmt"
	headers "./helpers/headers"
)

func main() {

	headers := headers.get()

	fmt.Println(headers.title)
}
