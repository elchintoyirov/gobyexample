// Go [muntazam ifodalar](https://en.wikipedia.org/wiki/Regular_expression) uchun o'rnatilgan qo'llab-quvvatlashni taklif etadi.
// Mana Go'da regexp bilan bog'liq keng tarqalgan
// vazifalarning ba'zi misollari.

package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	// Bu pattern satrga mos kelishini tekshiradi.
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// Yuqorida biz satr patternini to'g'ridan-to'g'ri
	// ishlatdik, lekin boshqa regexp vazifalari uchun
	// optimallashtirilgan `Regexp` structini `Compile` qilishingiz
	// kerak bo'ladi.
	r, _ := regexp.Compile("p([a-z]+)ch")

	// Bu structlarda ko'plab metodlar mavjud. Mana ilgari
	// ko'rganimizga o'xshash moslik testi.
	fmt.Println(r.MatchString("peach"))

	// Bu regexp uchun moslikni topadi.
	fmt.Println(r.FindString("peach punch"))

	// Bu ham birinchi moslikni topadi, lekin mos kelgan
	// matn o'rniga moslikning boshlanish va tugash
	// indekslarini qaytaradi.
	fmt.Println("idx:", r.FindStringIndex("peach punch"))

	// `Submatch` variantlari ham butun pattern mosliklari,
	// ham o'sha mosliklar ichidagi submosliklar haqida
	// ma'lumotni o'z ichiga oladi. Masalan, bu `p([a-z]+)ch`
	// va `([a-z]+)` ikkalasi uchun ma'lumot qaytaradi.
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// Xuddi shunday, bu mosliklar va submosliklar
	// indekslari haqida ma'lumot qaytaradi.
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// Bu funksiyalarning `All` variantlari faqat birinchisiga
	// emas, kirishdagi barcha mosliklarga tatbiq etiladi.
	// Masalan, regexp uchun barcha mosliklarni topish.
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// Bu `All` variantlari yuqorida ko'rgan boshqa
	// funksiyalar uchun ham mavjud.
	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	// Bu funksiyalarga ikkinchi argument sifatida manfiy
	// bo'lmagan butun son berish mosliklar sonini cheklaydi.
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// Yuqoridagi misollarimizda satr argumentlari bor edi va
	// `MatchString` kabi nomlar ishlatdik. Biz `[]byte`
	// argumentlarini ham bera olamiz va funksiya nomidan
	// `String`ni olib tashlay olamiz.
	fmt.Println(r.Match([]byte("peach")))

	// Muntazam ifodalar bilan global o'zgaruvchilar
	// yaratganda `Compile`ning `MustCompile` variantidan
	// foydalanishingiz mumkin. `MustCompile` xato qaytarish
	// o'rniga panika qiladi, bu uni global o'zgaruvchilar
	// uchun ishlatishni xavfsizroq qiladi.
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	// `regexp` paketi satrlarning qismlarini boshqa qiymatlar
	// bilan almashtirish uchun ham ishlatilishi mumkin.
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// `Func` varianti mos kelgan matnni berilgan funksiya
	// bilan o'zgartirishga imkon beradi.
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
