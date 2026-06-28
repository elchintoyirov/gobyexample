// Go'da holatni boshqarishning asosiy mexanizmi kanallar
// orqali aloqa qilishdir. Buni biz, masalan,
// [worker pools](worker-pools) misolida ko'rgan edik. Ammo
// holatni boshqarishning yana bir nechta usullari mavjud.
// Bu yerda biz bir nechta goroutinalar tomonidan
// foydalaniladigan _atomik hisoblagichlar_ uchun
// `sync/atomic` paketidan foydalanishni ko'rib chiqamiz.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// (Har doim musbat bo'lgan) hisoblagichimizni ifodalash
	// uchun atomik butun son tipidan foydalanamiz.
	var ops atomic.Uint64

	// WaitGroup barcha goroutinalar o'z ishlarini
	// tugatishini kutishimizga yordam beradi.
	var wg sync.WaitGroup

	// Biz 50 ta goroutina ishga tushiramiz, ularning har biri
	// hisoblagichni aniq 1000 marta oshiradi.
	for range 50 {
		wg.Go(func() {
			for range 1000 {
				// Hisoblagichni atomik tarzda oshirish uchun `Add`dan foydalanamiz.
				ops.Add(1)
			}
		})
	}

	// Barcha goroutinalar tugaguncha kutamiz.
	wg.Wait()

	// Bu yerda hech qaysi goroutina 'ops'ga yozmayapti, lekin
	// `Load` yordamida boshqa goroutinalar uni (atomik tarzda)
	// yangilab turgan paytda ham qiymatni atomik tarzda
	// xavfsiz o'qish mumkin.
	fmt.Println("ops:", ops.Load())
}
