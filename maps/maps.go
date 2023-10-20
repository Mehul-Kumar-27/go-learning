package maps

import "fmt"

func MapsInGo() {
	fmt.Println("Hello To Maps")

	mapOfLanguages := make(map[string]int)

	mapOfLanguages["key1"] = 1
	mapOfLanguages["key2"] = 2

	fmt.Println("The Value of mapOfLanguages is ", mapOfLanguages)
	fmt.Println("The Value of mapOfLanguages is ", mapOfLanguages["key3"])
}
