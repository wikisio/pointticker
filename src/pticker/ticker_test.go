package pticker

import (
	"fmt"
	"testing"
	"time"
)

func TestPTicker(t *testing.T) {
	ticker := NewTicker(5 * time.Second)
	defer ticker.Stop()

	ticker.Start()

	done := make(chan struct{})
	go func() {
		time.Sleep(22 * time.Second)
		done <- struct{}{}
	}()

	fmt.Println(time.Now())
	for {
		select {
		case m := <-ticker.C:
			fmt.Println(m)
		case <-done:
			return
		}
	}
}
