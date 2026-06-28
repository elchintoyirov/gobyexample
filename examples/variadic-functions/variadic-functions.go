// [_Variadik (o'zgaruvchan argumentli) funksiyalar_](https://en.wikipedia.org/wiki/Variadic_function)
// istalgan miqdordagi oxirgi argumentlar bilan chaqirilishi
// mumkin. Masalan, `fmt.Println` keng tarqalgan variadik
// funksiyadir.

package main

import "fmt"

// Quyida argument sifatida ixtiyoriy miqdordagi `int`
// larni qabul qiladigan funksiya keltirilgan.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	// Funksiya ichida `nums` ning tipi `[]int` ga
	// teng. Biz `len(nums)` ni chaqirishimiz, uni `range`
	// bilan aylanib chiqishimiz va hokazo qilishimiz mumkin.
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	// Variadik funksiyalarni odatdagi tarzda alohida
	// argumentlar bilan chaqirish mumkin.
	sum(1, 2)
	sum(1, 2, 3)

	// Agar sizda allaqachon slice ichida bir nechta
	// argument bo'lsa, ularni variadik funksiyaga
	// shu tarzda `func(slice...)` orqali qo'llang.
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
