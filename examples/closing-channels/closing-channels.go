// Kanalni _yopish_ unga boshqa qiymatlar jo'natilmasligini
// bildiradi. Bu kanalning qabul qiluvchilariga ishning
// tugaganini xabar qilish uchun foydali bo'lishi mumkin.

package main

import "fmt"

// Bu misolda biz bajariladigan ishni `main()`
// goroutinasidan worker goroutinasiga yetkazish uchun
// `jobs` kanalidan foydalanamiz. Worker uchun boshqa
// ishlar qolmaganda, biz `jobs` kanalini `close` qilamiz.
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// Mana worker goroutinasi. U `j, more := <-jobs` orqali
	// `jobs`dan takror-takror qabul qiladi. Qabul qilishning
	// ushbu maxsus 2 qiymatli shaklida, agar `jobs` `close`
	// qilingan va kanaldagi barcha qiymatlar allaqachon
	// qabul qilingan bo'lsa, `more` qiymati `false` bo'ladi.
	// Biz buni barcha ishlarni bajarib bo'lganimizda `done`
	// orqali xabar berish uchun ishlatamiz.
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// Bu `jobs` kanali orqali workerga 3 ta ish jo'natadi,
	// so'ngra uni yopadi.
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// Biz workerni avval ko'rgan
	// [sinxronizatsiya](channel-synchronization) yondashuvi
	// yordamida kutamiz.
	<-done

	// Yopilgan kanaldan o'qish darhol muvaffaqiyatli amalga
	// oshadi va asosiy tipning nol qiymatini qaytaradi.
	// Ixtiyoriy ikkinchi qaytariladigan qiymat, agar qabul
	// qilingan qiymat kanalga muvaffaqiyatli jo'natish amali
	// orqali yetkazilgan bo'lsa `true`, yoki kanal yopiq va
	// bo'sh bo'lgani uchun hosil qilingan nol qiymat bo'lsa
	// `false` bo'ladi.
	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}
