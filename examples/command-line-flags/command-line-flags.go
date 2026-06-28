// [_Buyruq qatori bayroqlari_](https://en.wikipedia.org/wiki/Command-line_interface#Command-line_option)
// buyruq qatori dasturlari uchun parametrlarni ko'rsatishning
// keng tarqalgan usulidir. Masalan, `wc -l`da `-l` buyruq
// qatori bayrog'idir.

package main

// Go asosiy buyruq qatori bayroqlarini tahlil qilishni
// qo'llab-quvvatlaydigan `flag` paketini taqdim etadi. Biz
// ushbu paketdan misol buyruq qatori dasturimizni amalga
// oshirish uchun foydalanamiz.
import (
	"flag"
	"fmt"
)

func main() {

	// Asosiy bayroq e'lonlari string, butun son va boolean
	// parametrlar uchun mavjud. Bu yerda biz `"foo"` standart
	// qiymati va qisqacha tavsifga ega `word` string
	// bayrog'ini e'lon qilamiz. Ushbu `flag.String` funksiyasi
	// string ko'rsatkichini qaytaradi (string qiymatini emas);
	// bu ko'rsatkichdan qanday foydalanishni quyida ko'ramiz.
	wordPtr := flag.String("word", "foo", "a string")

	// Bu `word` bayrog'iga o'xshash yondashuvdan foydalanib,
	// `numb` va `fork` bayroqlarini e'lon qiladi.
	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	// Shuningdek, dasturning boshqa joyida e'lon qilingan
	// mavjud o'zgaruvchidan foydalanadigan parametrni e'lon
	// qilish ham mumkin. E'tibor bering, bayroq e'lon qilish
	// funksiyasiga ko'rsatkichni uzatishimiz kerak.
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// Barcha bayroqlar e'lon qilingach, buyruq qatorini
	// tahlil qilishni bajarish uchun `flag.Parse()`ni
	// chaqiring.
	flag.Parse()

	// Bu yerda biz tahlil qilingan parametrlarni va
	// oxiridagi pozitsion argumentlarni chiqaramiz. E'tibor
	// bering, haqiqiy parametr qiymatlarini olish uchun
	// ko'rsatkichlarni masalan `*wordPtr` ko'rinishida
	// dereferensiya qilishimiz kerak.
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
