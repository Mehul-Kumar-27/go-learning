package waitgroupexample

import (
	"fmt"
	"sync"
	"time"
)

func getUserDetails(userId int, userChan chan string, wg *sync.WaitGroup) {
	time.Sleep(50 * time.Millisecond)

	userChan <- "We have the user Details"
	wg.Done()
}

func getUserRecommendations(userId int, userChan chan string, wg *sync.WaitGroup) {
	time.Sleep(100 * time.Millisecond)
	userChan <- ("We have the user recommendations")
	wg.Done()
}

func getUserLikes(userId int, userChan chan string, wg *sync.WaitGroup) {
	time.Sleep(150 * time.Millisecond)

	userChan <- ("We Have the user Likes")
	wg.Done()
}

func WaitGroupExample() {
	userId := 10
	now := time.Now()

	userChan := make(chan string, 3)
	wg := &sync.WaitGroup{}

	wg.Add(3)
	go getUserDetails(userId, userChan, wg)
	go getUserRecommendations(userId, userChan, wg)
	go getUserLikes(userId, userChan, wg)

	wg.Wait()
	close(userChan)

	for userResponse := range userChan {
		fmt.Println(userResponse)
	}

	fmt.Println(time.Since(now))

}
