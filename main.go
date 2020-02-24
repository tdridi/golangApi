package main

import (
	"fmt"
)

// Get Actors state
type Actors struct {
	number    int
	actorName string
}

func main() {
	//Array Example
	inspection := [3]int{5, 6, 89}
	fmt.Printf("Inspections: %v\n", inspection)

	//Slice Example
	files := []int{5, 6, 89}
	fmt.Println(files)

	//Map example
	provincePopulation := make(map[string]int)
	provincePopulation = map[string]int{
		"Quebec":           12000000,
		"Ontario":          13000000,
		"Manitoba":         1700000,
		"British-Columbia": 230988776487494,
		"Saskatchewan":     60549549854,
		"Alberta":          98409384098098,
	}
	fmt.Printf("Province Plopulation : %v\n", provincePopulation)

	//Struct Examples
	firstActor := Actors{
		number:    5145318616,
		actorName: "Talel Dridi",
	}

	secondActor := struct {
		age  int
		name string
	}{45, "Kalamata"}
	fmt.Printf("Fist Actor : %v\n", firstActor.actorName)
	fmt.Printf("Second Actor : %v\n", secondActor.name)

}
