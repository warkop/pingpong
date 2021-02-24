package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	c <- sum // send sum to c
}

func lessonOne() {
	go say("world")
	say("hello")
}

func lessonTwo() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
	fmt.Println(":len(s)/2", s[:len(s)/2])
	fmt.Println("len(s)/2:", s[len(s)/2:])
}

func main() {
	lessonTwo()
}
