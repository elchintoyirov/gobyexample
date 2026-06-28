// Bu misolda goroutinalar va kanallar yordamida
// _ishchi hovuzini_ qanday amalga oshirishni ko'rib chiqamiz.

package main

import (
	"fmt"
	"time"
)

// Quyida ishchi, uning bir nechta parallel nusxalarini
// ishga tushiramiz. Bu ishchilar `jobs` kanalida ish
// qabul qiladi va mos keladigan natijalarni `results` ga
// yuboradi. Resurs talab qiladigan vazifani simulyatsiya
// qilish uchun har bir ish uchun bir soniya uxlaymiz.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// Ishchilar hovuzimizdan foydalanish uchun ularga ish
	// yuborishimiz va natijalarini yig'ishimiz kerak. Buning
	// uchun 2 ta kanal yaratamiz.
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Bu 3 ta ishchini ishga tushiradi, ular hozircha
	// ish yo'qligi sababli dastlab bloklangan bo'ladi.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Bu yerda 5 ta `jobs` yuboramiz va keyin ishlarimiz
	// shu bilan tugaganini bildirish uchun o'sha kanalni
	// `close` qilamiz.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Nihoyat ishning barcha natijalarini yig'amiz. Bu
	// shuningdek ishchi goroutinalar tugaganligini ham
	// ta'minlaydi. Bir nechta goroutinani kutishning
	// muqobil usuli [WaitGroup](waitgroups) dan
	// foydalanishdir.
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
