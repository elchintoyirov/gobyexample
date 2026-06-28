// Kanallarni funksiya parametrlari sifatida ishlatganda,
// kanal faqat qiymat jo'natish yoki qabul qilish uchun
// mo'ljallanganligini ko'rsatishingiz mumkin. Bu aniqlik
// dasturning tip xavfsizligini oshiradi.

package main

import "fmt"

// Bu `ping` funksiyasi faqat qiymat jo'natish uchun kanalni
// qabul qiladi. Bu kanaldan qabul qilishga urinish
// kompilyatsiya vaqtidagi xato bo'lardi.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// `pong` funksiyasi qabul qilish uchun bitta kanalni
// (`pings`) va jo'natish uchun ikkinchisini (`pongs`) qabul
// qiladi.
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
