package main

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World!")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld() // asynchronous : tidak menunggu sampai selesai
	fmt.Println("Upps")

	time.Sleep(1 * time.Second) // seolah-olah synchronous : menunggu sampai selesai semua baru func test berhenti
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 10000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}