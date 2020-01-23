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
	leftCS, rightCS *ChopS
}

// Philosopher Eat Method
func (p Philo) eat() {
	for {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("eating")

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
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5]}
	}

	// Start the Dining
	for i := 0; i < 5; i++ {
		go philos[i].eat()
	}
}
