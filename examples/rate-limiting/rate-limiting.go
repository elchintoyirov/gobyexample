// [_So'rovlarni cheklash_](https://en.wikipedia.org/wiki/Rate_limiting)
// resurslardan foydalanishni nazorat qilish va xizmat
// sifatini saqlash uchun muhim mexanizmdir. Go
// goroutinalar, kanallar va [tickerlar](tickers) yordamida
// so'rovlarni cheklashni nafis tarzda qo'llab-quvvatlaydi.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Avval so'rovlarni cheklashning asosiy holatini ko'rib
	// chiqamiz. Faraz qilaylik, kelayotgan so'rovlarni qayta
	// ishlashni cheklamoqchimiz. Bu so'rovlarni xuddi shu
	// nomdagi kanal orqali xizmat ko'rsatamiz.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Bu `limiter` kanali har 200 millisekundda bitta qiymat
	// qabul qiladi. Bu bizning so'rovlarni cheklash
	// sxemamizdagi regulyatordir.
	limiter := time.Tick(200 * time.Millisecond)

	// Har bir so'rovga xizmat ko'rsatishdan oldin `limiter`
	// kanalidan qabul qilishni kutib bloklanish orqali, biz
	// o'zimizni har 200 millisekundda 1 so'rov bilan
	// cheklaymiz.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// Umumiy cheklovni saqlagan holda, so'rovlarni cheklash
	// sxemamizda qisqa muddatli so'rovlar to'lqinlariga ruxsat
	// berishni xohlashimiz mumkin. Buni limiter kanalimizni
	// buferlash orqali amalga oshira olamiz. Bu `burstyLimiter`
	// kanali 3 tagacha hodisaning to'lqiniga ruxsat beradi.
	burstyLimiter := make(chan time.Time, 3)

	// Ruxsat etilgan to'lqinni ifodalash uchun kanalni to'ldiramiz.
	for range 3 {
		burstyLimiter <- time.Now()
	}

	// Har 200 millisekundda `burstyLimiter`ga, uning 3 lik
	// chegarasigacha, yangi qiymat qo'shishga harakat qilamiz.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Endi yana 5 ta kelayotgan so'rovni simulyatsiya qilamiz.
	// Ulardan dastlabki 3 tasi `burstyLimiter`ning to'lqin
	// imkoniyatidan foyda ko'radi.
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
