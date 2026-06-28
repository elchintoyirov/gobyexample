// _Funksiyalar_ Go'da markaziy o'rin tutadi. Biz
// funksiyalarni bir nechta turli misollar bilan o'rganamiz.

package main

import "fmt"

// Mana ikkita `int` qabul qilib, ularning yig'indisini
// `int` sifatida qaytaradigan funksiya.
func plus(a int, b int) int {

	// Go aniq `return` talab qiladi, ya'ni u oxirgi
	// ifodaning qiymatini avtomatik ravishda qaytarmaydi.
	return a + b
}

// Bir xil tipdagi bir nechta ketma-ket parametringiz
// bo'lsa, tipni e'lon qiladigan oxirgi parametrgacha bir
// xil tipdagi parametrlar uchun tip nomini tushirib
// qoldirishingiz mumkin.
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	// Funksiyani kutganingizdek, `name(args)` bilan
	// chaqiring.
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
