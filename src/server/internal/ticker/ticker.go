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

var summaryResult Summary

type SummaryService struct{}

type SummaryInt interface {
	IncreaseTotalUnits()
	// ResetMessagePerSecond()
	// IncreaseTotalReached()
	// GetSummary()
	Tick()
}

func (s *SummaryService) Tick() {

	fmt.Println("tick triggered")
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Total units reached: ", summaryResult.totalReached)
				fmt.Println("Total units processed: ", summaryResult.totalUnits)
				return
			case t := <-ticker.C:
				fmt.Println("Messages per second at ", t, "are: ", summaryResult.messagePerSecond)
			}
		}
	}()

	time.Sleep(30 * time.Second)
	ticker.Stop()
	done <- true
}

func (s *SummaryService) IncreaseTotalUnits() {
	fmt.Println("Increasing", summaryResult.totalUnits)
	summaryResult.mu.Lock()
	defer summaryResult.mu.Unlock()
	summaryResult.totalUnits++
	summaryResult.messagePerSecond++
	fmt.Println("Increased", summaryResult.totalUnits)
}

//  func (s *Summary) IncreaseTotalReached() {
//  	s.mu.Lock()
//  	defer s.mu.Unlock()
//  	s.totalReached++
// / 	s.messagePerSecond++
// // }

// func (s *Summary) ResetMessagePerSecond() {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()
// 	s.messagePerSecond = 0
// }

// func (s *Summary) GetSummary() (*Summary, error) {
// 	return s, nil
// }
