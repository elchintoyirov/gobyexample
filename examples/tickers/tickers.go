// [Timerlar](timers) kelajakda biror narsani bir marta qilishni
// xohlaganingizda kerak bo'ladi - _tickerlar_ esa biror narsani
// muntazam oraliqlarda takror-takror qilishni xohlaganingizda
// kerak bo'ladi. Mana biz to'xtatmagunimizcha vaqti-vaqti bilan
// "tiqillab" turuvchi tickerga misol.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Tickerlar timerlarga o'xshash mexanizmdan foydalanadi:
	// qiymatlar yuboriladigan kanal. Bu yerda biz har 500ms da
	// kelayotgan qiymatlarni kutish uchun kanalda o'rnatilgan
	// `select` dan foydalanamiz.
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// Tickerlarni timerlar kabi to'xtatish mumkin. Ticker
	// to'xtatilgandan keyin u o'z kanalida boshqa qiymatlarni
	// qabul qilmaydi. Biz o'zimiznikini 1600ms dan keyin
	// to'xtatamiz.
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
