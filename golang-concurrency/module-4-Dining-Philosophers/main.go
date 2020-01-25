package main

import (
	"fmt"
	"sync"
)

// ChopS is Chopsticks structure
type ChopS struct {
	sync.Mutex
}

// Philo is Philosophers structure
type Philo struct {
	number          int
	leftCS, rightCS *ChopS
}

// Philosopher Eat Method
func (p Philo) eat(sem chan struct{}) {
	sem <- struct{}{}
	defer func() { <-sem }()
	for i := 0; i < 3; i++ {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("starting to eat %d\n", p.number)
		fmt.Printf("finishing eating %d\n", p.number)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
}

func main() {
	// Initialization
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i + 1, CSticks[i], CSticks[(i+1)%5]}
	}

	sem := make(chan struct{}, 2)

	var wg sync.WaitGroup
	// Start the Dining
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			philos[i].eat(sem)
		}(i)
	}
	wg.Wait()
}
