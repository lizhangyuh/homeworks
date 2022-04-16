package main

import (
	"testing"
	"time"
)

func Test10Sum(t *testing.T) {
	bc := NewBucketCounter(1*1000, 10)
	for i := 0; i < 10; i++ {
		bc.Accept(1)
		time.Sleep(1 * time.Second)
	}

	if sum := bc.Sum(); sum != 10 {
		t.Errorf("sum expected to be 10, but %d got", sum)
	}
}

func Test15Sum(t *testing.T) {
	bc := NewBucketCounter(1*1000, 10)
	for i := 0; i < 15; i++ {
		bc.Accept(1)
		time.Sleep(1 * time.Second)
	}

	if sum := bc.Sum(); sum != 10 {
		t.Errorf("sum expected to be 10, but %d got", sum)
	}
}

func TestFastSum(t *testing.T) {
	bc := NewBucketCounter(1*1000, 10)
	for i := 0; i < 15; i++ {
		bc.Accept(1)
		time.Sleep(100 * time.Millisecond)
	}
	time.Sleep(10 * time.Second)
	for i := 0; i < 10; i++ {
		bc.Accept(1)
		time.Sleep(100 * time.Millisecond)
	}

	if sum := bc.Sum(); sum != 10 {
		t.Errorf("sum expected to be 10, but %d got", sum)
	}
}

func TestSlowSum(t *testing.T) {
	bc := NewBucketCounter(1*1000, 10)
	for i := 0; i < 5; i++ {
		bc.Accept(1)
		time.Sleep(4 * time.Second)
	}

	if sum := bc.Sum(); sum != 3 {
		t.Errorf("sum expected to be 3, but %d got", sum)
	}
}

func TestConcurrency(t *testing.T) {
	bc := NewBucketCounter(1*1000, 10)
	stop := make(chan struct{})

	go func() {
		for i := 0; i < 5; i++ {
			bc.Accept(1)
			time.Sleep(1 * time.Second)
		}
		stop <- struct{}{}
	}()

	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		bc.Accept(1)
	}

	<-stop

	if sum := bc.Sum(); sum != 10 {
		t.Errorf("sum expected to be 10, but %d got", sum)
	}
}
