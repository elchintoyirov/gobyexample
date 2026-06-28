// Bir nechta goroutinaning tugashini kutish uchun biz
// *wait group* (kutish guruhi) dan foydalanishimiz mumkin.

package main

import (
	"fmt"
	"sync"
	"time"
)

// Bu har bir goroutinada ishga tushiradigan funksiyamiz.
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	// Resurs talab qiladigan vazifani simulyatsiya qilish uchun uxlaymiz.
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	// Bu WaitGroup shu yerda ishga tushirilgan barcha
	// goroutinalarning tugashini kutish uchun ishlatiladi.
	// Eslatma: agar WaitGroup funksiyalarga aniq uzatilsa,
	// bu *ko'rsatkich orqali* amalga oshirilishi kerak.
	var wg sync.WaitGroup

	// `WaitGroup.Go` yordamida bir nechta goroutina ishga tushiramiz
	for i := 1; i <= 5; i++ {
		wg.Go(func() {
			worker(i)
		})
	}

	// `wg` tomonidan boshlangan barcha goroutinalar
	// tugaguncha bloklaymiz. Goroutina o'zi chaqirgan
	// funksiya qaytarganda tugagan hisoblanadi.
	wg.Wait()

	// E'tibor bering, bu yondashuvda ishchilardan
	// xatolarni tarqatishning oddiy usuli yo'q. Yanada
	// murakkab holatlar uchun
	// [errgroup paketidan](https://pkg.go.dev/golang.org/x/sync/errgroup)
	// foydalanishni ko'rib chiqing.
}
