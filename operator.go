package main

import "fmt"

func main() {
	const name string = "golang"
	fmt.Printf("nama saya %s", name)

	var values = (5 + 2) * 3
	fmt.Println(values)

	var conditionOne bool = 3 > 5                  //false
	var conditionTwo bool = 4 > 2                  // true
	var conditionThree bool = "golang" == "GOLANG" // false
	var conditionFour bool = 10 != 3.8             // true
	var conditionFive = 11 <= 11

	fmt.Println("Condition One", conditionOne)
	fmt.Println("Condition Two", conditionTwo)
	fmt.Println("Condition Three", conditionThree)
	fmt.Println("Condition Four", conditionFour)
	fmt.Println("Condition Five", conditionFive)

	var right = true
	var wrong = false

	var wrongAndRight = wrong && right //false
	fmt.Printf("wrong  && right \t(%t) \n", wrongAndRight)

	var wrongOrRight = wrong || right //true
	fmt.Printf("wrong || right \t(%t) \n", wrongOrRight)

	var wrongReverse = !wrong // true
	fmt.Printf("!wrong \t(%t)\n", wrongReverse)
}
