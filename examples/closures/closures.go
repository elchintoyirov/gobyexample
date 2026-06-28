// Go [_anonim funksiyalar_](https://en.wikipedia.org/wiki/Anonymous_function)ni
// qo'llab-quvvatlaydi, ular <a href="https://en.wikipedia.org/wiki/Closure_(computer_science)"><em>closure</em></a>larni hosil qilishi mumkin.
// Anonim funksiyalar funksiyani nomlamasdan, joyida e'lon
// qilmoqchi bo'lganingizda foydalidir.

package main

import "fmt"

// Bu `intSeq` funksiyasi boshqa funksiyani qaytaradi, biz
// uni `intSeq`ning tanasida anonim tarzda aniqlaymiz.
// Qaytarilgan funksiya closure hosil qilish uchun `i`
// o'zgaruvchisini _qamrab oladi_ (closes over).
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// Biz `intSeq`ni chaqiramiz va natijani (funksiyani)
	// `nextInt`ga o'zlashtiramiz. Ushbu funksiya qiymati o'z
	// `i` qiymatini qamrab oladi, u har gal `nextInt`ni
	// chaqirganimizda yangilanadi.
	nextInt := intSeq()

	// `nextInt`ni bir necha marta chaqirib, closure
	// ta'sirini ko'ring.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// Holat o'sha aniq funksiyaga xos ekanligini tasdiqlash
	// uchun, yangisini yaratib sinab ko'ring.
	newInts := intSeq()
	fmt.Println(newInts())
}
