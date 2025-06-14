package belajar_golang_goroutines

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