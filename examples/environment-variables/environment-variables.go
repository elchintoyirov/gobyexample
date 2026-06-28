// [Muhit o'zgaruvchilari](https://en.wikipedia.org/wiki/Environment_variable)
// [Unix dasturlariga konfiguratsiya ma'lumotlarini
// yetkazish](https://www.12factor.net/config) uchun universal mexanizmdir.
// Keling, muhit o'zgaruvchilarini qanday o'rnatish, olish va ro'yxatlashni ko'rib chiqaylik.

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// Kalit/qiymat juftligini o'rnatish uchun `os.Setenv` dan
	// foydalaning. Kalit uchun qiymatni olish uchun `os.Getenv` dan
	// foydalaning. Agar kalit muhitda mavjud bo'lmasa, bu bo'sh satrni
	// qaytaradi.
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// Muhitdagi barcha kalit/qiymat juftliklarini ro'yxatlash uchun
	// `os.Environ` dan foydalaning. Bu `KEY=value` ko'rinishidagi satrlar
	// slice'ini qaytaradi. Kalit va qiymatni olish uchun ularni
	// `strings.SplitN` qilishingiz mumkin. Bu yerda biz barcha kalitlarni chop etamiz.
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
