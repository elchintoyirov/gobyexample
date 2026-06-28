// _Slice_lar Go'da muhim ma'lumotlar tipi bo'lib,
// ketma-ketliklarga massivlardan ko'ra kuchliroq interfeys
// beradi.

package main

import (
	"fmt"
	"slices"
)

func main() {

	// Massivlardan farqli o'laroq, slice'lar faqat o'zlari
	// saqlaydigan elementlar bilan tiplanadi (elementlar soni
	// bilan emas). Initsializatsiya qilinmagan slice nilga teng
	// va uzunligi 0.
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// Nolga teng bo'lmagan uzunlikdagi slice yaratish uchun
	// o'rnatilgan `make`dan foydalaning. Bu yerda biz uzunligi
	// `3` bo'lgan `string`lar slice'ini yaratamiz (dastlab nol
	// qiymatli). Standart holatda yangi slice'ning sig'imi uning
	// uzunligiga teng; agar slice oldindan o'sishini bilsak,
	// `make`ga qo'shimcha parametr sifatida sig'imni oshkora
	// berish mumkin.
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// Massivlardagi kabi qiymatlarni o'rnatish va olishimiz mumkin.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// `len` kutilganidek slice uzunligini qaytaradi.
	fmt.Println("len:", len(s))

	// Bu asosiy operatsiyalardan tashqari, slice'lar ularni
	// massivlardan boyroq qiladigan yana bir nechtasini
	// qo'llab-quvvatlaydi. Ulardan biri o'rnatilgan `append`
	// bo'lib, u bir yoki bir nechta yangi qiymatni o'z ichiga
	// olgan slice qaytaradi. E'tibor bering, yangi slice qiymati
	// olishimiz mumkinligi sababli `append`dan qaytadigan
	// qiymatni qabul qilishimiz kerak.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slice'lar `copy` ham qilinishi mumkin. Bu yerda biz `s`
	// bilan bir xil uzunlikdagi bo'sh `c` slice'ini yaratamiz va
	// `s`dan `c`ga nusxalaymiz.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slice'lar `slice[low:high]` sintaksisi bilan "slice"
	// operatorini qo'llab-quvvatlaydi. Masalan, bu `s[2]`,
	// `s[3]` va `s[4]` elementlarining slice'ini oladi.
	l := s[2:5]
	fmt.Println("sl1:", l)

	// Bu `s[5]`gacha (lekin uni o'z ichiga olmagan holda) slice qiladi.
	l = s[:5]
	fmt.Println("sl2:", l)

	// Va bu `s[2]`dan (uni o'z ichiga olgan holda) slice qiladi.
	l = s[2:]
	fmt.Println("sl3:", l)

	// Slice uchun o'zgaruvchini bir qatorda e'lon qilib,
	// initsializatsiya ham qila olamiz.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// `slices` paketi slice'lar uchun bir qancha foydali
	// yordamchi funksiyalarni o'z ichiga oladi.
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// Slice'lar ko'p o'lchovli ma'lumotlar tuzilmalariga
	// birlashtirilishi mumkin. Ko'p o'lchovli massivlardan
	// farqli o'laroq, ichki slice'larning uzunligi har xil
	// bo'lishi mumkin.
	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
