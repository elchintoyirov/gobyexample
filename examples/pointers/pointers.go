// Go <em><a href="https://en.wikipedia.org/wiki/Pointer_(computer_programming)">ko'rsatkichlar</a></em> ni qo'llab-quvvatlaydi,
// bu dasturingiz ichidagi qiymatlar va yozuvlarga havolalarni
// uzatish imkonini beradi.

package main

import "fmt"

// Biz ko'rsatkichlar qiymatlarga nisbatan qanday ishlashini
// 2 ta funksiya bilan ko'rsatamiz: `zeroval` va `zeroptr`.
// `zeroval` `int` parametriga ega, shuning uchun argumentlar
// unga qiymat bo'yicha uzatiladi. `zeroval` chaqiruvchi
// funksiyadagidan farqli `ival` ning nusxasini oladi.
func zeroval(ival int) {
	ival = 0
}

// `zeroptr` esa aksincha `*int` parametriga ega, ya'ni u
// `int` ko'rsatkichini qabul qiladi. Funksiya tanasidagi
// `*iptr` kodi keyin ko'rsatkichni uning xotira manzilidan
// o'sha manzildagi joriy qiymatga _dereferensiya qiladi_.
// Dereferensiya qilingan ko'rsatkichga qiymat o'zlashtirish
// havola qilingan manzildagi qiymatni o'zgartiradi.
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// `&i` sintaksisi `i` ning xotira manzilini, ya'ni `i` ga
	// ko'rsatkichni beradi.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// Ko'rsatkichlarni ham chop etish mumkin.
	fmt.Println("pointer:", &i)
}
