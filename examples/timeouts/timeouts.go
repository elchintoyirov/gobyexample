// _Vaqt tugashi_ tashqi resurslarga ulanadigan yoki
// boshqacha sababga ko'ra bajarilish vaqtini cheklashi
// kerak bo'lgan dasturlar uchun muhim. Go da vaqt
// tugashini amalga oshirish kanallar va `select`
// tufayli oson va nafis.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Misol uchun, 2s dan keyin natijasini `c1` kanaliga
	// qaytaradigan tashqi chaqiruvni bajaryapmiz deylik.
	// E'tibor bering, kanal buferlangan, shuning uchun
	// goroutinadagi yuborish bloklamaydi. Bu kanal hech
	// qachon o'qilmagan taqdirda goroutina oqib ketishining
	// oldini olish uchun keng tarqalgan usul.
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// Quyida vaqt tugashini amalga oshiruvchi `select`.
	// `res := <-c1` natijani kutadi, `<-time.After` esa 1s
	// vaqt tugashidan keyin yuboriladigan qiymatni kutadi.
	// `select` tayyor bo'lgan birinchi qabul bilan davom
	// etgani uchun, operatsiya ruxsat etilgan 1s dan ko'proq
	// vaqt olsa, vaqt tugashi holatini olamiz.
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// Agar 3s lik uzunroq vaqt tugashiga ruxsat bersak,
	// `c2` dan qabul qilish muvaffaqiyatli bo'ladi va
	// natijani chop etamiz.
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
