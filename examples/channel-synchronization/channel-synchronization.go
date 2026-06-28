// Kanallardan goroutinalar o'rtasidagi bajarilishni
// sinxronlash uchun foydalanishimiz mumkin. Mana
// goroutinaning tugashini kutish uchun bloklovchi qabul
// qilishdan foydalanish misoli. Bir nechta goroutinaning
// tugashini kutayotganingizda, [WaitGroup](waitgroups)dan
// foydalanishni afzal ko'rishingiz mumkin.

package main

import (
	"fmt"
	"time"
)

// Bu biz goroutinada ishga tushiradigan funksiya. `done`
// kanali ushbu funksiyaning ishi tugaganini boshqa
// goroutinaga xabar berish uchun ishlatiladi.
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// Ishimiz tugaganini bildirish uchun qiymat jo'natamiz.
	done <- true
}

func main() {

	// Worker goroutinasini ishga tushiramiz va unga xabar
	// berish uchun kanalni beramiz.
	done := make(chan bool, 1)
	go worker(done)

	// Kanal orqali workerdan xabar olgunimizcha bloklaymiz.
	<-done
}
