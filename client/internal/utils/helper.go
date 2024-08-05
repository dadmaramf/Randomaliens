package utils

import (
	"log"
	"time"
)

func StartTicker(done chan bool, stat *Stat) {
	ticker := time.NewTicker(time.Second * 10)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				log.Println(stat.GetStats())
			}
		}

	}()

}
