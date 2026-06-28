// Go _bir nechta qaytariladigan qiymatlar_ ni ichki o'rnatilgan
// holda qo'llab-quvvatlaydi. Bu xususiyat idiomatik Go da tez-tez
// ishlatiladi, masalan funksiyadan ham natija ham xato qiymatlarini
// qaytarish uchun.

package main

import "fmt"

// Ushbu funksiya imzosidagi `(int, int)` funksiya 2 ta `int`
// qaytarishini ko'rsatadi.
func vals() (int, int) {
	return 3, 7
}

func main() {

	// Bu yerda biz chaqiruvdan kelgan 2 ta turli qaytariladigan
	// qiymatni _bir nechta o'zlashtirish_ bilan ishlatamiz.
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// Agar sizga qaytarilgan qiymatlarning faqat bir qismi kerak
	// bo'lsa, bo'sh identifikator `_` dan foydalaning.
	_, c := vals()
	fmt.Println(c)
}
