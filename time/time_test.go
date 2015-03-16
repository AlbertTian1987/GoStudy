package time

import (
	"fmt"
	"testing"
	"time"
)

func Test_TimeChan(t *testing.T) {

	result := make(chan int)
	go func(ch chan int) {
		time.Sleep(3 * time.Second)
		ch <- 4
	}(result)
	for {
		select {
		case <-time.After(5 * time.Second):
			fmt.Println("time out")
		case <-result:
			fmt.Println(result)
		}
	}

}
