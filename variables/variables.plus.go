package main

import "fmt"

func main() {
	// var declaration plus
	var (
		hero               = "Goku"
		enemyOne, enemyTwo = "Vegeta", "Cell"
	)
	fmt.Println("name for Hero: "+hero, " - name of enemy one: "+enemyOne, " - name of enemy two: "+enemyTwo)
}
