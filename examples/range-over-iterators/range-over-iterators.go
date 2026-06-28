// 1.23 versiyasidan boshlab, Go
// [iteratorlar](https://go.dev/blog/range-functions) uchun
// qo'llab-quvvatlash qo'shdi, bu bizga deyarli har qanday
// narsa bo'ylab range qilish imkonini beradi!

package main

import (
	"fmt"
	"iter"
	"slices"
	"strings"
)

// Keling, [oldingi misol](generics) dagi `List` tipini yana
// ko'rib chiqaylik. O'sha misolda ro'yxatdagi barcha
// elementlarning slice ini qaytaradigan `AllElements`
// metodimiz bor edi. Go iteratorlari bilan buni yaxshiroq
// bajarishimiz mumkin - quyida ko'rsatilganidek.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// All _iterator_ ni qaytaradi, bu Go da
// [maxsus imzoga](https://pkg.go.dev/iter#Seq) ega funksiyadir.
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		// Iterator funksiyasi parametr sifatida boshqa
		// funksiyani qabul qiladi, kelishuvga ko'ra `yield` deb
		// ataladi (lekin nom ixtiyoriy bo'lishi mumkin). U biz
		// takrorlamoqchi bo'lgan har bir element uchun `yield`
		// ni chaqiradi va ehtimoliy erta to'xtatish uchun
		// `yield` ning qaytariladigan qiymatini hisobga oladi.
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

// Takrorlash uchun asosiy ma'lumotlar strukturasi talab
// etilmaydi va hatto chekli bo'lishi ham shart emas! Mana
// Fibonachchi sonlari bo'ylab iterator qaytaradigan funksiya:
// u `yield` `true` qaytarishda davom etgan ekan, ishlashda
// davom etadi.
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 0, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	// `List.All` iterator qaytargani uchun, biz undan oddiy
	// `range` siklida foydalanishimiz mumkin.
	for e := range lst.All() {
		fmt.Println(e)
	}

	// [slices](https://pkg.go.dev/slices) kabi paketlarda
	// iteratorlar bilan ishlash uchun bir qancha foydali
	// funksiyalar mavjud. Masalan, `Collect` istalgan
	// iteratorni qabul qiladi va uning barcha qiymatlarini
	// slice ga to'playdi.
	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	// Standart kutubxona paketlari endi iterator yordamchilarini
	// ham taqdim etadi. Masalan, `strings.SplitSeq` avval
	// natija slice ini qurmasdan bayt slice ining qismlari
	// bo'ylab takrorlanadi.
	for part := range strings.SplitSeq("go-by-example", "-") {
		fmt.Printf("part: %s\n", part)
	}

	for n := range genFib() {

		// Sikl `break` ga yoki erta qaytishga duch kelishi
		// bilan, iteratorga uzatilgan `yield` funksiyasi
		// `false` qaytaradi.
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}
