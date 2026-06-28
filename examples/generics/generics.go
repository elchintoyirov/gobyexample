// 1.18-versiyadan boshlab Go _generiklar_ uchun qo'llab-quvvatlash
// qo'shdi, ular _tip parametrlari_ deb ham ataladi.

package main

import "fmt"

// Generik funksiyaga misol sifatida, `SlicesIndex` istalgan
// `comparable` tipdagi slice va o'sha tipdagi elementni
// qabul qiladi hamda v ning s dagi birinchi uchrash indeksini,
// agar mavjud bo'lmasa -1 ni qaytaradi. `comparable`
// cheklovi shuni anglatadiki, biz ushbu tipdagi qiymatlarni
// `==` va `!=` operatorlari bilan solishtira olamiz. Ushbu
// tip imzosining batafsilroq tushuntirishi uchun [shu blog
// postiga](https://go.dev/blog/deconstructing-type-parameters)
// qarang. E'tibor bering, bu funksiya standart kutubxonada
// [slices.Index](https://pkg.go.dev/slices#Index) sifatida
// mavjud.
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

// Generik tipga misol sifatida, `List` istalgan tipdagi
// qiymatlarga ega bir tomonlama bog'langan ro'yxatdir.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// Biz generik tiplarda xuddi oddiy tiplardagi kabi
// metodlarni aniqlay olamiz, lekin tip parametrlarini
// joyida saqlab qolishimiz kerak. Tip `List` emas, balki
// `List[T]` dir.
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// AllElements barcha List elementlarini slice sifatida
// qaytaradi. Keyingi misolda biz maxsus tiplarning barcha
// elementlari ustidan iteratsiya qilishning yanada idiomatik
// usulini ko'ramiz.
func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var s = []string{"foo", "bar", "zoo"}

	// Generik funksiyalarni chaqirganda biz ko'pincha _tipni
	// avtomatik aniqlash_ ga tayanishimiz mumkin. E'tibor
	// bering, `SlicesIndex` ni chaqirganda `S` va `E` tiplarini
	// ko'rsatishimiz shart emas - kompilyator ularni avtomatik
	// ravishda aniqlaydi.
	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	// ... garchi biz ularni aniq ham ko'rsatishimiz mumkin.
	_ = SlicesIndex[[]string, string](s, "zoo")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())
}
