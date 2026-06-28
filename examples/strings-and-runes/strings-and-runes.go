// Go satri faqat o'qish uchun mo'ljallangan baytlar slice'idir.
// Til va standart kutubxona satrlarni alohida tarzda -
// [UTF-8](https://en.wikipedia.org/wiki/UTF-8) da kodlangan
// matn konteynerlari sifatida ko'rib chiqadi. Boshqa tillarda
// satrlar "belgilar"dan tashkil topadi. Go'da belgi tushunchasi
// `rune` deb ataladi - bu Unicode kod nuqtasini ifodalovchi
// butun son. [Ushbu Go blog posti](https://go.dev/blog/strings)
// mavzuga yaxshi kirish bo'ladi.

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// `s` - bu tay tilida "salom" so'zini ifodalovchi literal
	// qiymat berilgan `string`. Go satr literallari UTF-8 da
	// kodlangan matndir.
	const s = "สวัสดี"

	// Satrlar `[]byte` ga ekvivalent bo'lgani uchun, bu uning
	// ichida saqlangan xom baytlar uzunligini beradi.
	fmt.Println("Len:", len(s))

	// Satrni indekslash har bir indeksdagi xom bayt qiymatlarini
	// beradi. Bu sikl `s` dagi kod nuqtalarini tashkil etuvchi
	// barcha baytlarning o'n oltilik qiymatlarini hosil qiladi.
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// Satrda nechta _runa_ borligini sanash uchun biz `utf8`
	// paketidan foydalanishimiz mumkin. E'tibor bering,
	// `RuneCountInString` ning ishlash vaqti satr o'lchamiga
	// bog'liq, chunki u har bir UTF-8 runani ketma-ket
	// dekodlashi kerak. Ba'zi tay belgilari bir nechta baytga
	// cho'zilishi mumkin bo'lgan UTF-8 kod nuqtalari bilan
	// ifodalanadi, shuning uchun bu sanoq natijasi kutilmagan
	// bo'lishi mumkin.
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// `range` sikli satrlarni alohida tarzda ishlaydi va har bir
	// `rune` ni uning satrdagi siljishi bilan birga dekodlaydi.
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// Biz `utf8.DecodeRuneInString` funksiyasidan oshkora
	// foydalanib xuddi shu iteratsiyaga erishishimiz mumkin.
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		// Bu funksiyaga `rune` qiymatini uzatishni namoyish etadi.
		examineRune(runeValue)
	}
}

func examineRune(r rune) {

	// Bittalik tirnoq ichiga olingan qiymatlar _runa literallari_
	// hisoblanadi. Biz `rune` qiymatini runa literali bilan
	// to'g'ridan-to'g'ri solishtirishimiz mumkin.
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
