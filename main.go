package main

import (
	"fmt"
	"math/rand"
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
}

func blockingChannel() {
	c := make(chan bool)
	c <- true
	fmt.Println("This line will never be printed")
}

func blockingChannelCont() {
	c := make(chan bool)
	go func() {
		<-c
	}()
	c <- true
	fmt.Println("This line is printed")
}

func bufferedChannel() {
	c := make(chan bool, 5)
	c <- true
	fmt.Println("This line is printed")
}

func anonymous() {
	for _, v := range []int{1, 2, 3, 4, 5} {
		go func(i int) {
			fmt.Println(i)
		}(v)
	}
}

func selectExample() {
	c := make(chan int)

	go func() {
		<-time.After(time.Duration(rand.Intn(2)) * time.Second)
		c <- 10
	}()

	select {
	case val := <-c:
		fmt.Println(val)
	case <-time.After(time.Duration(rand.Intn(2)) * time.Second):
		fmt.Println("timeout")
	}
}

func defaultSelection() {
	tick := time.Tick(100 * time.Millisecond)
	fmt.Println("tick:", tick)
	boom := time.Tick(500 * time.Millisecond)
	fmt.Println("boom:", boom)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
			fmt.Println("<-tick: ", <-tick)
		case <-boom:
			fmt.Println("BOOM!")
			fmt.Println("<-boom: ", <-boom)
			return
		default:
			fmt.Println("   .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func main() {
	defaultSelection()
}
