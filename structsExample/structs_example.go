package structs_example

import "fmt"

func StructsExample() {
	person := User{"Mehul"}

	fmt.Println(person)

	person.ChangeName()

	fmt.Println(person)

}

type User struct {
	Name string
}

func (u *User) ChangeName() {
	u.Name = "Mehul Kumar"

	fmt.Println(u.Name)

}
