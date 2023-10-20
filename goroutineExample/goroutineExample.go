package goroutineExample

import (
	"fmt"
	"time"
)

func changeTheNumber(a *int, name string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 100)
		currentTime := time.Now().Format("2006-01-02 15:04:05.000000")
		fmt.Printf("The current time is %v in %v iteration %v\n", currentTime, name, i)
		fmt.Printf("The value of a is %v and currently run by %v in thr iteration %v \n", *a, name, i)
		*a = *a + *a
	}
}

func GoRoutineExample() {
	a := 1
	go changeTheNumber(&a, "Go")
	changeTheNumber(&a, "Simple")
	time.Sleep(time.Millisecond * 100)
}
