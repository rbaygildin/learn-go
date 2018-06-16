package main

import (
	"testing"
	"fmt"
	"time"
)

func TestGoRoutines(t *testing.T){
	f := func(name string) {
		fmt.Println("Hello,", name, "!")
	}
	defer f("Last")
	f("James")
	go f("Doe")
	f("Peter")
}

func TestChannels(t *testing.T){
	ch := make(chan string)
	f := func(channel chan string) {
		channel <- "Hello world\n"
	}
	go f(ch)
	message := <- ch
	fmt.Print(message)
}

func TestBufferedChannels(t *testing.T) {
	ch := make(chan string, 2)
	f := func(channel chan string) {
		channel <- "Hello world\n"
		channel <- "Are you awesome?\n"
	}
	go f(ch)
	message := <- ch
	fmt.Print(message)
	message = <- ch
	fmt.Print(message)
}

func TestChannelsTimeOut(t *testing.T){
	ch := make(chan string)
	f := func(channel chan string) {
		time.Sleep(2 * time.Second)
		channel <- "Hello world\n"
	}
	go f(ch)
	var message string
	select {
		case message = <- ch:
			fmt.Println(message)
		case <- time.After(1 * time.Second):
			fmt.Println("Time out\n")
	}
}