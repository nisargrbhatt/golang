package main

import (
	"fmt"
	"math/rand"
	"time"
)

func process(c chan int) {
	c <- 10
}

func send5Items(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}

}

func basicExample() {
	var c = make(chan int)

	go process(c)
	fmt.Println(<-c)

	go send5Items(c)
	for i := range c {
		fmt.Println(i)
	}

}

var MAX_PANEER_PRICE float32 = 5
var MAX_CHEESE_PRICE float32 = 10

func checkPaneerPrices(website string, c chan string) {
	for {
		time.Sleep(time.Second * 1)
		var paneerPrice = rand.Float32() * 20
		if paneerPrice < MAX_PANEER_PRICE {
			c <- website
			break
		}
	}
}

func checkCheesePrices(website string, c chan string) {
	for {
		time.Sleep(time.Second * 2)
		var cheesePrice = rand.Float32() * 20
		if cheesePrice < MAX_CHEESE_PRICE {
			c <- website
			break
		}
	}
}

func sendMessage(paneerChan chan string, cheeseChan chan string) {
	select {
	case website := <-paneerChan:
		fmt.Println("Found paneer deal from", website)
	case website := <-cheeseChan:
		fmt.Println("Found cheese deal from", website)
	}
}

func selectChanExample() {
	var paneerChan = make(chan string)
	var cheeseChan = make(chan string)

	var wesites = []string{"google", "amazon", "ebay", "walmart"}
	for _, website := range wesites {
		go checkPaneerPrices(website, paneerChan)
		go checkCheesePrices(website, cheeseChan)
	}
	sendMessage(paneerChan, cheeseChan)
}

func main() {
	// basicExample()
	selectChanExample()
}
