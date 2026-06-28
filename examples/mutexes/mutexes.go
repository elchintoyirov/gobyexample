// Oldingi misolda biz [atomik amallar](atomic-counters)
// yordamida oddiy hisoblagich holatini qanday boshqarishni
// ko'rdik. Murakkabroq holat uchun bir nechta goroutina
// orasida ma'lumotlarga xavfsiz kirish uchun
// [_mutex_](https://en.wikipedia.org/wiki/Mutual_exclusion) dan
// foydalanishimiz mumkin.

package main

import (
	"fmt"
	"sync"
)

// Container hisoblagichlar map ini saqlaydi; biz uni bir
// nechta goroutinadan parallel yangilamoqchi bo'lganimiz uchun
// kirishni sinxronlash maqsadida `Mutex` qo'shamiz.
// E'tibor bering, mutexlar nusxalanmasligi kerak, shuning
// uchun bu `struct` uzatilsa, u ko'rsatkich orqali
// bajarilishi kerak.
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	// `counters` ga kirishdan oldin mutex ni bloklang; uni
	// funksiya oxirida [defer](defer) operatori yordamida
	// blokdan chiqaring.
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := Container{
		// E'tibor bering, mutex ning nol qiymati o'zicha ishlatsa
		// bo'ladi, shuning uchun bu yerda ishga tushirish talab etilmaydi.
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// Bu funksiya nomlangan hisoblagichni siklda oshiradi.
	doIncrement := func(name string, n int) {
		for range n {
			c.inc(name)
		}
	}

	// Bir nechta goroutinani parallel ishga tushiramiz; e'tibor
	// bering, ularning barchasi bir xil `Container` ga kiradi,
	// va ulardan ikkitasi bir xil hisoblagichga kiradi.
	wg.Go(func() {
		doIncrement("a", 10000)
	})

	wg.Go(func() {
		doIncrement("a", 10000)
	})

	wg.Go(func() {
		doIncrement("b", 10000)
	})

	// Goroutinalar tugashini kutamiz
	wg.Wait()
	fmt.Println(c.counters)
}
