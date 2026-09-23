package game

import (
	"sync"
	"time"
)

type Clock struct {
	CurrentTime time.Time
	Mu          sync.RWMutex
}

func (c *Clock) BackgroundTime() {
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		c.CurrentTime = c.CurrentTime.Add(1 * time.Second)
	}
}

type StartTime struct {
	StartTime time.Time
}
