// Go'ning `slices` paketi o'rnatilgan va foydalanuvchi
// aniqlagan tiplar uchun saralashni amalga oshiradi. Avval
// o'rnatilgan tiplar uchun saralashni ko'rib chiqamiz.

package main

import (
	"fmt"
	"slices"
)

func main() {

	// Saralash funksiyalari generikdir va har qanday _tartibli_
	// o'rnatilgan tip uchun ishlaydi. Tartibli tiplar ro'yxati
	// uchun [cmp.Ordered](https://pkg.go.dev/cmp#Ordered)ga qarang.
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	// `int`larni saralash misoli.
	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:   ", ints)

	// Slice allaqachon saralangan tartibda ekanligini
	// tekshirish uchun ham `slices` paketidan foydalanishimiz
	// mumkin.
	s := slices.IsSorted(ints)
	fmt.Println("Sorted: ", s)
}
