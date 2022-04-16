package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	bc := NewBucketCounter(1*1000, 10)
	go func() {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for {
			time.Sleep(time.Duration(r.Intn(20) * 1000 * int(time.Millisecond)))
			bc.Accept(1)
		}
	}()

	for {
		time.Sleep(10 * 1000 * time.Millisecond)
		fmt.Printf("total count: %d\n", bc.Sum())
	}
}
