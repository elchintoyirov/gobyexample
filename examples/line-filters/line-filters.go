// _Qator filtri_ - bu stdin'dan kiritishni o'qiydigan,
// uni qayta ishlaydigan va so'ngra ba'zi olingan natijani
// stdout'ga chiqaradigan dasturlarning keng tarqalgan
// turidir. `grep` va `sed` keng tarqalgan qator
// filtrlaridir.

// Mana Go'da barcha kiritilgan matnning bosh harfli
// versiyasini yozadigan qator filtri misoli. Siz ushbu
// shablondan o'zingizning Go qator filtrlaringizni yozish
// uchun foydalanishingiz mumkin.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Buferlanmagan `os.Stdin` ni buferlangan skaner bilan
	// o'rash bizga skanerni keyingi tokenga o'tkazadigan
	// qulay `Scan` metodini beradi; bu standart skanerda
	// keyingi qatordir.
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// `Text` kiritishdan joriy tokenni, bu yerda keyingi
		// qatorni qaytaradi.
		ucl := strings.ToUpper(scanner.Text())

		// Bosh harfli qatorni chiqarish.
		fmt.Println(ucl)
	}

	// `Scan` davomida xatolarni tekshirish. Fayl oxiri
	// kutiladi va `Scan` tomonidan xato sifatida xabar
	// qilinmaydi.
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
