// _Goroutina_ - bu yengil bajarilish oqimidir.

package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// Faraz qilaylik, bizda `f(s)` funksiya chaqiruvi bor.
	// Mana uni odatdagidek, sinxron tarzda ishga tushirib
	// chaqiramiz.
	f("direct")

	// Ushbu funksiyani goroutinada chaqirish uchun
	// `go f(s)` dan foydalaning. Bu yangi goroutina
	// chaqiruvchi goroutina bilan parallel bajariladi.
	go f("goroutine")

	// Shuningdek, anonim funksiya chaqiruvi uchun goroutina
	// ishga tushirishingiz mumkin.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Endi bizning ikkita funksiya chaqiruvimiz alohida
	// goroutinalarda asinxron ishlamoqda. Ularning tugashini
	// kuting (yanada ishonchli yondashuv uchun
	// [WaitGroup](waitgroups) dan foydalaning).
	time.Sleep(time.Second)
	fmt.Println("done")
}
