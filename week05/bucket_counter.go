package main

import (
	"fmt"
	"sync"
	"time"
)

type Event struct {
	Count int64
}

func (event *Event) String() string {
	return fmt.Sprintf("%d", event.Count)
}

func (event *Event) Reset(count int64) {
	event.Count = count
}

func (event *Event) Add(count int64) {
	event.Count += count
}

type BucketCounter struct {
	StartAt      int64
	LastOffset   int64
	LastIndex    int64
	Mutex        *sync.RWMutex
	DurationInMs int64
	BucketNum    int64
	Events       []*Event
}

func (bc *BucketCounter) Accept(input int64) {
	now := time.Now().UnixMilli()
	bc.Mutex.Lock()
	defer bc.Mutex.Unlock()

	offset := (now - bc.StartAt) / bc.DurationInMs
	index := offset % bc.BucketNum
	fmt.Printf("index: %d\n", index)

	if offset-bc.LastOffset >= bc.BucketNum {
		// 超时清空窗口数据
		for _, e := range bc.Events {
			e.Reset(0)
		}
	} else {
		// 距离上次accept之间窗口数据置0
		if index > bc.LastIndex {
			for i, e := range bc.Events {
				if i > int(bc.LastIndex) && i <= int(index) {
					e.Reset(0)
				}
			}
		}

		if index < bc.LastIndex {
			for i, e := range bc.Events {
				if i <= int(index) || i > int(bc.LastIndex) {
					e.Reset(0)
				}
			}
		}
	}

	bc.Events[index].Add(input)
	fmt.Printf("offset: %d, event %d add %d\n", offset, index, input)

	bc.LastIndex = index
	bc.LastOffset = offset

	fmt.Printf("events: %v\n", bc.Events)
}

func (bc *BucketCounter) Sum() int64 {
	sum := int64(0)

	bc.Mutex.RLock()
	defer bc.Mutex.RUnlock()

	for _, e := range bc.Events {
		sum += e.Count
	}

	return sum
}

func NewBucketCounter(durationInMs int64, num int64) *BucketCounter {
	counter := &BucketCounter{
		StartAt:      time.Now().UnixMilli(),
		LastIndex:    0,
		LastOffset:   0,
		DurationInMs: durationInMs,
		BucketNum:    num,
		Events:       make([]*Event, num),
		Mutex:        &sync.RWMutex{},
	}

	for i := range counter.Events {
		counter.Events[i] = &Event{Count: 0}
	}

	return counter
}
