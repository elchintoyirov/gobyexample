// Ba'zan to'plamni uning tabiiy tartibidan boshqa narsa
// bo'yicha saralashni xohlaymiz. Masalan, faraz qilaylik,
// satrlarni alifbo tartibida emas, balki uzunligi bo'yicha
// saralashni xohladik. Mana Go'da maxsus saralashga misol.

package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"peach", "banana", "kiwi"}

	// Satr uzunliklari uchun taqqoslash funksiyasini amalga
	// oshiramiz. `cmp.Compare` bunga yordam beradi.
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	// Endi `fruits`ni nom uzunligi bo'yicha saralash uchun bu
	// maxsus taqqoslash funksiyasi bilan `slices.SortFunc`ni
	// chaqira olamiz.
	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	// O'rnatilgan tip bo'lmagan qiymatlar slice'ini saralash
	// uchun xuddi shu usuldan foydalanishimiz mumkin.
	type Person struct {
		name string
		age  int
	}

	people := []Person{
		Person{name: "Jax", age: 37},
		Person{name: "TJ", age: 25},
		Person{name: "Alex", age: 72},
	}

	// `people`ni yoshi bo'yicha `slices.SortFunc` yordamida
	// saralang.
	//
	// Eslatma: agar `Person` struct katta bo'lsa, slice
	// o'rniga `*Person`ni o'z ichiga olishini xohlashingiz va
	// saralash funksiyasini shunga moslab sozlashingiz mumkin.
	// Shubhangiz bo'lsa, [benchmark](testing-and-benchmarking)
	// qiling!
	slices.SortFunc(people,
		func(a, b Person) int {
			return cmp.Compare(a.age, b.age)
		})
	fmt.Println(people)
}
