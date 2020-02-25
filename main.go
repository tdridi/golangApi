package main

import (
	"fmt"
)

//Get Actors state
type Actors struct {
	number    int
	actorName string
}

/** Get Actors speciality */
type Speciality struct {
	Actors
	firstSpeciality  string
	secondSpeciality []string
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
	fmt.Printf("Province Plopulation Quebec : %v\n", provincePopulation["Quebec"])

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

	thirdActor := Speciality{}
	thirdActor.actorName = "felaini"
	thirdActor.number = 42
	thirdActor.firstSpeciality = "Horror"
	thirdActor.secondSpeciality = append(thirdActor.secondSpeciality, "Drama")

	fmt.Printf("Third Actor: %v", thirdActor)
}
