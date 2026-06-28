// Go'ning _select_i bir nechta kanal operatsiyalarini
// kutishga imkon beradi. Goroutinalar va kanallarni select
// bilan birlashtirish Go'ning kuchli xususiyatidir.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Misolimiz uchun ikkita kanal bo'ylab select qilamiz.
	c1 := make(chan string)
	c2 := make(chan string)

	// Har bir kanal biroz vaqtdan so'ng qiymat qabul qiladi,
	// bu masalan parallel goroutinalarda bajarilayotgan
	// bloklovchi RPC operatsiyalarini simulyatsiya qilish uchun.
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// Bu qiymatlarning ikkalasini bir vaqtning o'zida kutish
	// uchun `select`dan foydalanamiz, har birini kelishi bilan
	// chop etamiz.
	for range 2 {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
