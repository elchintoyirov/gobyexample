// Go'da _massiv_ — bu aniq uzunlikka ega bo'lgan,
// raqamlangan elementlar ketma-ketligidir. Odatdagi Go
// kodida [slices](slices) ancha keng tarqalgan; massivlar
// esa ba'zi maxsus holatlarda foydali bo'ladi.

package main

import "fmt"

func main() {

	// Bu yerda biz aniq 5 ta `int` saqlaydigan `a` massivini
	// yaratamiz. Elementlarning tipi ham, uzunligi ham
	// massiv tipining bir qismidir. Standart holatda massiv
	// nol qiymatga ega bo'ladi, bu `int`lar uchun `0`larni
	// bildiradi.
	var a [5]int
	fmt.Println("emp:", a)

	// `array[index] = value` sintaksisi yordamida indeksdagi
	// qiymatni o'rnatishimiz, `array[index]` orqali esa
	// qiymatni olishimiz mumkin.
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// O'rnatilgan `len` funksiyasi massivning uzunligini
	// qaytaradi.
	fmt.Println("len:", len(a))

	// Massivni bitta qatorda e'lon qilish va initsializatsiya
	// qilish uchun ushbu sintaksisdan foydalaning.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Shuningdek, `...` yordamida elementlar sonini
	// kompilyatorning o'zi hisoblashiga ham yo'l qo'yishingiz
	// mumkin.
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Agar indeksni `:` bilan ko'rsatsangiz, oradagi
	// elementlar nolga tenglashtiriladi.
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	// Massiv tiplari bir o'lchamlidir, lekin tiplarni
	// birlashtirib, ko'p o'lchamli ma'lumotlar tuzilmalarini
	// qurishingiz mumkin.
	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// Ko'p o'lchamli massivlarni bir vaqtning o'zida yaratish
	// va initsializatsiya qilish ham mumkin.
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
