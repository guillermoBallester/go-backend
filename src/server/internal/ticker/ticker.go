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

type SummaryResponse struct {
	totalUnits   int `json:"totalUnits"`
	totalReached int `json:"totalReached"`
}

var summaryResult Summary

type SummaryService struct{}

type SummaryInt interface {
	IncreaseTotalUnits()
	ResetMessagePerSecond()
	IncreaseTotalReached()
	GetSummary() (SummaryResponse, error)
	Tick()
}

func (s *SummaryService) Tick() {

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

func (s *SummaryService) IncreaseTotalReached() {
	summaryResult.mu.Lock()
	defer summaryResult.mu.Unlock()
	summaryResult.totalReached++
	summaryResult.messagePerSecond++
}

func (s *SummaryService) ResetMessagePerSecond() {
	summaryResult.mu.Lock()
	defer summaryResult.mu.Unlock()
	summaryResult.messagePerSecond = 0
}

func (s *SummaryService) GetSummary() (SummaryResponse, error) {
	summaryResult.mu.Lock()
	defer summaryResult.mu.Unlock()
	return SummaryResponse{
		totalUnits:   summaryResult.totalUnits,
		totalReached: summaryResult.totalReached,
	}, nil
}
