// _range_ turli ichki o'rnatilgan ma'lumotlar
// strukturalaridagi elementlar bo'ylab takrorlanadi. Keling,
// `range` ni biz allaqachon o'rgangan ba'zi ma'lumotlar
// strukturalari bilan qanday ishlatishni ko'rib chiqamiz.

package main

import "fmt"

func main() {

	// Bu yerda biz `range` dan slice dagi sonlarni yig'ish uchun
	// foydalanamiz. Massivlar ham xuddi shunday ishlaydi.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// Massivlar va slice lardagi `range` har bir element uchun
	// ham indeks ham qiymatni beradi. Yuqorida bizga indeks
	// kerak emas edi, shuning uchun uni bo'sh identifikator
	// `_` bilan e'tiborsiz qoldirdik. Biroq, ba'zan bizga
	// indekslar haqiqatan kerak bo'ladi.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// Map dagi `range` kalit/qiymat juftliklari bo'ylab takrorlanadi.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// `range` shuningdek map ning faqat kalitlari bo'ylab takrorlanishi mumkin.
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// Satrlardagi `range` Unicode kod nuqtalari bo'ylab
	// takrorlanadi. Birinchi qiymat `rune` ning boshlang'ich
	// bayt indeksi, ikkinchisi esa `rune` ning o'zi. Batafsil
	// ma'lumot uchun [Strings and Runes](strings-and-runes) ga qarang.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
