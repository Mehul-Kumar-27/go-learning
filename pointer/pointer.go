package pointer

import "fmt"

func PointerExample() {
	fmt.Println("Hello To Pointers")

	var ptr *int
	myNumber := 23

	ptr = &myNumber

	fmt.Println("The Value of Ptr is ", ptr)
	fmt.Println("The Value of Ptr is ", *ptr)
	 
}
