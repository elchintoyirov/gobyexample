// [Oldingi](range-over-built-in-types) misolda biz `for` va
// `range` asosiy ma'lumotlar strukturalari bo'ylab takrorlashni
// qanday ta'minlashini ko'rdik. Biz bu sintaksisdan kanaldan
// qabul qilingan qiymatlar bo'ylab takrorlash uchun ham
// foydalanishimiz mumkin.

package main

import "fmt"

func main() {

	// Biz `queue` kanalidagi 2 ta qiymat bo'ylab takrorlanamiz.
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// Bu `range` `queue` dan qabul qilingani sari har bir
	// element bo'ylab takrorlanadi. Biz yuqorida kanalni
	// `close` qilganimiz uchun, takrorlash 2 ta elementni
	// qabul qilgandan keyin to'xtaydi.
	for elem := range queue {
		fmt.Println(elem)
	}
}
