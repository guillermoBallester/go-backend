package ticker

import (
	"fmt"
	"sync"
	"time"
)

type Summary struct {
	totalUnits       int
	totalReached     int
	messagePerSecond int
	mu               sync.Mutex
}

func (s *Summary) Tick() {

	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Total units reached: ", s.totalReached)
				fmt.Println("Total units processed: ", s.totalUnits)
				return
			case t := <-ticker.C:
				fmt.Println("Messages per second at ", t, "are: ", s.messagePerSecond)
			}
			s.ResetMessagePerSecond()
		}

	}()

	time.Sleep(10 * time.Second)
	ticker.Stop()
	done <- true
}

func (s *Summary) IncreaseTotalUnits() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.totalUnits++
	s.messagePerSecond++
}

func (s *Summary) IncreaseTotalReached() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.totalReached++
	s.messagePerSecond++
}

func (s *Summary) ResetMessagePerSecond() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.messagePerSecond = 0
}
