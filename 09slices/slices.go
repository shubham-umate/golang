package main

import (
	"fmt"
	"sort"
)

func main() {

	var fruitList = []string{"Applse", "Peach", "Tomato"}
	fmt.Printf(" Type of fruitlist is %T\n", fruitList)
	fmt.Println(fruitList)
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 345
	highScore[2] = 456
	highScore[3] = 567
	// highScore[4] = 667

	highScore = append(highScore, 111, 222, 333)
	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	var courses = []string{"reactjs", "js", "c", "c++"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
