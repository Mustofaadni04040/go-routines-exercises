package main

import (
	"fmt"
	"time"
)

func task1() {
	for i := 0; i < 5; i++ {
		fmt.Println("task1:", i)
		time.Sleep(700 * time.Millisecond) // simulasikan proses lama
	}
}

func task2() {
	for i := range 5 {
		fmt.Println("task2:", i)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	go task1() // jika task1 lama maka akan lanjut ke task b begitu sebaliknya | Output akan campur aduk antar Task1 & 2
	go task2() // Ketika task1 sedang Sleep, Go scheduler bisa langsung beralih ke task2.

	time.Sleep(5 * time.Second) // tunggu semua selesai
}
