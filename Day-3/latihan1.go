package main

import (
	"fmt"
	"math/rand"
	"time"
)


// 	Buatlah sebuah program yang menentukan apakah seorang pahlawan dapat bertahan hidup atau tidak. 
// Alkisah ada seorang pahlawan yang ditugaskan pergi ke sebuah kastil di mana kastil tersebut
// dikelilingi oleh banyak naga. Seekor naga dapat dibunuh dengan 2 peluru. Permasalahannya adalah 
// pahlawan tersebut tidak tau berapa banyak naga yang ada dan berapa banyak peluru yang harus ia bawa. 
// Jika pahlawan tersebut berhasil bertahan hidup kembalikan nilai true begitu pula sebaliknya

func pahlawan(naga, peluru int) (sisaNaga int, statusPahlawan bool){
	if naga>(peluru/2){
		sisaNaga = naga-(peluru/2)
		statusPahlawan = false
	} else {
		if naga-(peluru/2)<=0{
			sisaNaga = 0
			statusPahlawan = true
		}else if naga-(peluru/2)>0{
			sisaNaga = naga-(peluru/2)
			statusPahlawan = true
		}
	}
	return
}

func main(){
	// Set the seed value to ensure random number generation
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 0 and 50
	randomNumber1 := rand.Intn(21)
	randomNumber2 := rand.Intn(51)

	naga := randomNumber1
	peluru := randomNumber2
	fmt.Printf("%d naga dan %d peluru\n\n", naga, peluru)

	sisaNaga := 0
	var statusPahlawan bool

	sisaNaga, statusPahlawan = pahlawan(naga, peluru)
	fmt.Println("Sisa naga =", sisaNaga)
	if statusPahlawan == true {
		fmt.Println("Pahlawan berhasil membasmi naga⚔️")
	} else {
		fmt.Println("Pahlawan telah gugur🪦")
	}
}