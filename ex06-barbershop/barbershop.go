package main

import (
	"fmt"
	"time"
)

func Barber(done chan int, room chan int) {
	customer := <-room
	fmt.Printf("Started working on a customer %v\n", customer)
	time.Sleep(time.Second * 5)
	done <- 1
	fmt.Printf("Finished working on a customer %v\n", customer)
}

func Customer(customer chan int, waitchan chan int) {
	for i := 1; i < 11; i++ {
		customer <- i
		time.Sleep(time.Second * 2)
	}
	for {
		select {
		case cust := <-waitchan:
			customer <- cust
			time.Sleep(time.Second * 15)
		}
	}
}

func main() {
	Is_sleeping := 1
	customer := make(chan int, 10)
	done := make(chan int, 10)
	room := make(chan int, 3)
	waitchan := make(chan int, 10)

	go Customer(customer, waitchan)

Selection:
	for {
		select {
		case cust := <-customer:
			if Is_sleeping == 1 {
				fmt.Println("Wake up!")
				Is_sleeping = 0
				room <- cust
				go Barber(done, room)
			} else if len(room) < 3 {
				room <- cust
				fmt.Printf("%v customer took a seat\n", cust)
			} else {
				fmt.Printf("%v customer has left the barbershop\n", cust)
				waitchan <- cust
			}
		case <-done:
			if len(room) == 0 {
				Is_sleeping = 1
				fmt.Println("Barber is sleeping")
				if len(waitchan) == 0 {
					fmt.Println("No customers for today. Good night!")
					break Selection
				}
			} else {
				go Barber(done, room)
			}
		}
	}
}
