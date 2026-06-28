// Standart holatda kanallar _buferlanmagan_ bo'ladi, ya'ni
// ular jo'natilgan qiymatni qabul qilishga tayyor mos
// qabul qiluvchi (`<- chan`) bo'lgan taqdirdagina
// jo'natishlarni (`chan <-`) qabul qiladi. _Buferlangan
// kanallar_ esa o'sha qiymatlar uchun mos qabul qiluvchisiz
// cheklangan miqdordagi qiymatlarni qabul qiladi.

package main

import "fmt"

func main() {

	// Bu yerda biz 2 tagacha qiymatni buferlaydigan satrlar
	// kanalini `make` qilamiz.
	messages := make(chan string, 2)

	// Bu kanal buferlangani uchun, ushbu qiymatlarni mos
	// parallel qabul qilishsiz kanalga jo'nata olamiz.
	messages <- "buffered"
	messages <- "channel"

	// Keyinchalik bu ikki qiymatni odatdagidek qabul qila
	// olamiz.
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
