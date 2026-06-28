// Go da _o'zgaruvchilar_ aniq e'lon qilinadi va kompilyator
// tomonidan, masalan, funksiya chaqiruvlarining tip
// to'g'riligini tekshirish uchun ishlatiladi.

package main

import "fmt"

func main() {

	// `var` 1 yoki undan ortiq o'zgaruvchilarni e'lon qiladi.
	var a = "initial"
	fmt.Println(a)

	// Bir nechta o'zgaruvchini bir vaqtda e'lon qilishingiz mumkin.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go ishga tushirilgan o'zgaruvchilarning tipini aniqlaydi.
	var d = true
	fmt.Println(d)

	// Mos keladigan ishga tushirishsiz e'lon qilingan
	// o'zgaruvchilar _nol qiymatli_ bo'ladi. Masalan, `int`
	// uchun nol qiymat `0` dir.
	var e int
	fmt.Println(e)

	// `:=` sintaksisi o'zgaruvchini e'lon qilish va ishga
	// tushirishning qisqartirilgan shaklidir, masalan bu
	// holatda `var f string = "apple"` uchun. Bu sintaksis
	// faqat funksiyalar ichida mavjud.
	f := "apple"
	fmt.Println(f)
}
