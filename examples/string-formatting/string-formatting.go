// Go `printf` an'anasida satrlarni formatlash uchun ajoyib
// qo'llab-quvvatlashni taklif qiladi. Mana satrlarni
// formatlashning ba'zi keng tarqalgan vazifalariga misollar.

package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	// Go umumiy Go qiymatlarini formatlash uchun mo'ljallangan
	// bir nechta chop etish "fe'l"larini taklif qiladi. Masalan,
	// bu bizning `point` struct'imizning nusxasini chop etadi.
	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p)

	// Agar qiymat struct bo'lsa, `%+v` varianti struct'ning
	// maydon nomlarini ham qo'shadi.
	fmt.Printf("struct2: %+v\n", p)

	// `%#v` varianti qiymatning Go sintaksisidagi ko'rinishini,
	// ya'ni o'sha qiymatni hosil qiladigan manba kod bo'lagini
	// chop etadi.
	fmt.Printf("struct3: %#v\n", p)

	// Qiymatning tipini chop etish uchun `%T` dan foydalaning.
	fmt.Printf("type: %T\n", p)

	// Mantiqiy (boolean) qiymatlarni formatlash oddiy.
	fmt.Printf("bool: %t\n", true)

	// Butun sonlarni formatlash uchun ko'plab variantlar mavjud.
	// Standart, 10 lik sanoq tizimidagi formatlash uchun `%d` dan
	// foydalaning.
	fmt.Printf("int: %d\n", 123)

	// Bu ikkilik (binary) ko'rinishni chop etadi.
	fmt.Printf("bin: %b\n", 14)

	// Bu berilgan butun songa mos keladigan belgini chop etadi.
	fmt.Printf("char: %c\n", 33)

	// `%x` o'n oltilik (hex) kodlashni ta'minlaydi.
	fmt.Printf("hex: %x\n", 456)

	// Kasr (float) sonlar uchun ham bir nechta formatlash
	// variantlari mavjud. Oddiy o'nlik formatlash uchun `%f` dan
	// foydalaning.
	fmt.Printf("float1: %f\n", 78.9)

	// `%e` va `%E` kasr sonni (biroz farqli ko'rinishlardagi)
	// ilmiy yozuvda formatlaydi.
	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)

	// Oddiy satr chop etish uchun `%s` dan foydalaning.
	fmt.Printf("str1: %s\n", "\"string\"")

	// Satrlarni Go manba kodidagi kabi qo'shtirnoq ichiga olish
	// uchun `%q` dan foydalaning.
	fmt.Printf("str2: %q\n", "\"string\"")

	// Avval ko'rilgan butun sonlardagi kabi, `%x` satrni 16 lik
	// sanoq tizimida, har bir kirish bayti uchun ikkita chiqish
	// belgisi bilan ko'rsatadi.
	fmt.Printf("str3: %x\n", "hex this")

	// Ko'rsatkichning ko'rinishini chop etish uchun `%p` dan
	// foydalaning.
	fmt.Printf("pointer: %p\n", &p)

	// Sonlarni formatlashda siz ko'pincha natijaviy raqamning
	// kengligi va aniqligini boshqarishni xohlaysiz. Butun
	// sonning kengligini belgilash uchun fe'ldagi `%` dan keyin
	// son qo'ying. Sukut bo'yicha natija o'ngga tekislanadi va
	// bo'sh joylar bilan to'ldiriladi.
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	// Siz chop etilgan kasr sonlarning kengligini ham belgilashingiz
	// mumkin, garchi odatda ayni paytda width.precision sintaksisi
	// bilan o'nlik aniqlikni ham cheklashni xohlaysiz.
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	// Chapga tekislash uchun `-` bayrog'idan foydalaning.
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// Satrlarni formatlashda ham kenglikni boshqarishni
	// xohlashingiz mumkin, ayniqsa ular jadvalga o'xshash
	// chiqishda tekislanishini ta'minlash uchun. Oddiy o'ngga
	// tekislangan kenglik uchun.
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	// Chapga tekislash uchun sonlardagi kabi `-` bayrog'idan
	// foydalaning.
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	// Hozirgacha biz formatlangan satrni `os.Stdout` ga chop
	// etadigan `Printf` ni ko'rdik. `Sprintf` satrni formatlaydi
	// va uni hech qaerga chop etmasdan qaytaradi.
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	// Siz `Fprintf` yordamida `os.Stdout` dan boshqa
	// `io.Writers` larga formatlab chop etishingiz mumkin.
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
