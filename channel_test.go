package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel) // close channel

	// mengirim data ke channel
	// channel <- "Ucok"

	// menerima data dari channel
	// data := <- channel

	// contoh jika mau mengirim langsung ke parameter func
	// fmt.Println(<-channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Ucok" // harus ada yang mengambil data nya jika tidak akan ke block kode dibawahnya
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <- channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)

	fmt.Println("selesai mengambil data dari channel")
}