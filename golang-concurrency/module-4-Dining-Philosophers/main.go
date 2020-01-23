// Refarence: http://rosettacode.org/wiki/Dining_philosophers#Mutexes_and_WaitGroup
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

// In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
var wg sync.WaitGroup

// Philosopher Eat Method
func (p Philo) eat() {
	for i := 0; i < 3; i++ {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("starting to eat %d\n", p.number)
		fmt.Printf("finishing eating %d\n", p.number)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
	wg.Done()
}

func main() {
	// Initialization
	wg.Add(2) // The host allows no more than 2 philosophers to eat concurrently.
	CStick0 := new(ChopS)
	CStickLeft := CStick0
	philos := make([]*Philo, 5)
	for i := 1; i < 5; i++ {
		CStickRight := new(ChopS)
		philos[i] = &Philo{i + 1, CStickLeft, CStickRight}
		CStickLeft = CStickRight
	}
	philos[0] = &Philo{1, CStick0, CStickLeft}

	// Start the Dining
	for i := 0; i < 5; i++ {
		go philos[i].eat()
	}
	wg.Wait()

}
