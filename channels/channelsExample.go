package channelsExample

import "fmt"

func ChannelsExample() {
	channelOne := make(chan string, 2)

	channelOne <- "1"
	channelOne <- "2"
	fmt.Println("Hello")
	fmt.Println(<-channelOne)
	channelOne <- "3"

	fmt.Println(<-channelOne)
}

type Payment struct {
	users    map[string]string
	UserChan chan string
}

func (p *Payment) Start() {
	fmt.Println("Start")
	p.loop()
}

func (p *Payment) loop() {

	for {
		fmt.Println("Loop")
		u := <-p.UserChan
		fmt.Println("Position 1")
		p.users[u] = u
		fmt.Println("Position 2")
	}
}

func NewPayment() *Payment {
	return &Payment{
		users: make(map[string]string),
	}
}

func (p *Payment) AddPayment(pay string) {
	p.users[pay] = pay
}
