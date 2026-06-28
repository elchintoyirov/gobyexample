// Go struct tiplarida aniqlangan _metod_larni qo'llab-quvvatlaydi.

package main

import "fmt"

type rect struct {
	width, height int
}

// Bu `area` metodi `*rect` _qabul qiluvchi tipiga_ ega.
func (r *rect) area() int {
	return r.width * r.height
}

// Metodlar ko'rsatkichli yoki qiymatli qabul qiluvchi tiplari
// uchun aniqlanishi mumkin. Mana qiymatli qabul qiluvchi misoli.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// Bu yerda biz struct imiz uchun aniqlangan 2 ta metodni chaqiramiz.
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go metod chaqiruvlari uchun qiymatlar va ko'rsatkichlar
	// o'rtasidagi o'tkazishni avtomatik boshqaradi. Metod
	// chaqiruvlarida nusxa olishdan qochish yoki metodga qabul
	// qiluvchi struct ni o'zgartirishga ruxsat berish uchun
	// ko'rsatkichli qabul qiluvchi tipidan foydalanishingiz mumkin.
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
