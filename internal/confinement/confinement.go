package confinement

import (
	"fmt"
	"sync"
)

func manageTicket(ticketChan <-chan int, done <-chan struct{}, tickets *int) {
	for {
		select {
		case userId := <-ticketChan:
			if *tickets > 0 {
				*tickets--
				fmt.Printf("%d purchased a ticket. Remaining tickets: %d\n", userId, *tickets)
			} else {
				fmt.Printf("%d found no tickets available.\n", userId)
			}
		case <-done:
			fmt.Printf("Tickets remaining: %d\n", *tickets)
		}
	}
}

func buyTicket(wg *sync.WaitGroup, ticketChan chan<- int, userId int) {
	defer wg.Done()
	ticketChan <- userId
}

func init() {
	var wg sync.WaitGroup
	tickets := 500
	ticketChan := make(chan int)
	doneChan := make(chan struct{})

	go manageTicket(ticketChan, doneChan, &tickets)

	for userId := 0; userId < 2000; userId++ {
		wg.Add(1)
		go buyTicket(&wg, ticketChan, userId)
	}

	wg.Wait()
	close(doneChan)
}
