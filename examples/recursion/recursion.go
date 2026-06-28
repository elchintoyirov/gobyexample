// Go
// <a href="https://en.wikipedia.org/wiki/Recursion_(computer_science)"><em>rekursiv funksiyalarni</em></a> qo'llab-quvvatlaydi.
// Mana klassik misol.

package main

import "fmt"

// Bu `fact` funksiyasi `fact(0)` bazaviy holatiga
// yetguncha o'zini chaqiradi.
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	// Anonim funksiyalar ham rekursiv bo'lishi mumkin, lekin
	// bu funksiya aniqlanishdan oldin uni saqlash uchun `var`
	// bilan o'zgaruvchini oshkora e'lon qilishni talab qiladi.
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		// `fib` ilgari `main`da e'lon qilinganligi sababli, Go
		// bu yerda `fib` bilan qaysi funksiyani chaqirishni
		// biladi.
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
