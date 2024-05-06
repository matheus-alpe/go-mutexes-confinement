package mutexesonly

import (
	"fmt"
	"sync"
)

func buyTicket(wg *sync.WaitGroup, mu *sync.Mutex, userId int, remainingTickets *int) {
	defer wg.Done()
	mu.Lock()
	if *remainingTickets > 0 {
		*remainingTickets--
		fmt.Printf("%d purchased a ticket. Remaining tickets: %d\n", userId, *remainingTickets)
	} else {

		fmt.Printf("%d found no tickets available.\n", userId)
	}
	mu.Unlock()
}

func init() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	tickets := 500

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go buyTicket(&wg, &mu, i, &tickets)
	}
	wg.Wait()
}
