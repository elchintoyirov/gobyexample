// _Map_lar Go'ning ichki o'rnatilgan [assotsiativ ma'lumotlar tipi](https://en.wikipedia.org/wiki/Associative_array)
// (boshqa tillarda ba'zan _hash_lar yoki _dict_lar deb ataladi).

package main

import (
	"fmt"
	"maps"
)

func main() {

	// Bo'sh map yaratish uchun ichki o'rnatilgan `make` dan foydalaning:
	// `make(map[key-type]val-type)`.
	m := make(map[string]int)

	// Kalit/qiymat juftliklarini odatiy `name[key] = val`
	// sintaksisi yordamida o'rnating.
	m["k1"] = 7
	m["k2"] = 13

	// Map ni masalan `fmt.Println` bilan chop etish uning barcha
	// kalit/qiymat juftliklarini ko'rsatadi.
	fmt.Println("map:", m)

	// Kalit uchun qiymatni `name[key]` orqali oling.
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// Agar kalit mavjud bo'lmasa, qiymat tipining
	// [nol qiymati](https://go.dev/ref/spec#The_zero_value)
	// qaytariladi.
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// Ichki o'rnatilgan `len` map ga qo'llanilganda kalit/qiymat
	// juftliklari sonini qaytaradi.
	fmt.Println("len:", len(m))

	// Ichki o'rnatilgan `delete` map dan kalit/qiymat
	// juftliklarini o'chiradi.
	delete(m, "k2")
	fmt.Println("map:", m)

	// Map dan *barcha* kalit/qiymat juftliklarini o'chirish uchun
	// `clear` ichki o'rnatilgan funksiyasidan foydalaning.
	clear(m)
	fmt.Println("map:", m)

	// Map dan qiymat olishda ixtiyoriy ikkinchi qaytariladigan
	// qiymat kalit map da mavjud bo'lganmi yoki yo'qligini
	// bildiradi. Bu yo'q kalitlarni `0` yoki `""` kabi nol
	// qiymatga ega kalitlardan ajratish uchun ishlatilishi
	// mumkin. Bu yerda bizga qiymatning o'zi kerak emas edi,
	// shuning uchun uni _bo'sh identifikator_ `_` bilan
	// e'tiborsiz qoldirdik.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// Shuningdek, ushbu sintaksis bilan bitta qatorda yangi map
	// e'lon qilishingiz va ishga tushirishingiz mumkin.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// `maps` paketi map lar uchun bir qancha foydali
	// yordamchi funksiyalarni o'z ichiga oladi.
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
