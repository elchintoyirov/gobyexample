// Go [base64
// kodlash/dekodlash](https://en.wikipedia.org/wiki/Base64)
// uchun o'rnatilgan qo'llab-quvvatlashni taqdim etadi.

package main

// Ushbu sintaksis `encoding/base64` paketini standart
// `base64` nomi o'rniga `b64` nomi bilan import qiladi. Bu
// quyida bizga biroz joy tejaydi.
import (
	b64 "encoding/base64"
	"fmt"
)

func main() {

	// Mana biz kodlaydigan/dekodlaydigan `string`.
	data := "abc123!?$*&()'-=@~"

	// Go ham standart, ham URL bilan mos keluvchi base64'ni
	// qo'llab-quvvatlaydi. Mana standart enkoder yordamida
	// qanday kodlash kerakligi. Enkoder `[]byte` talab
	// qiladi, shuning uchun `string`imizni o'sha tipga
	// o'tkazamiz.
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// Dekodlash xato qaytarishi mumkin, agar kirish
	// ma'lumotining to'g'ri shakllanganini oldindan
	// bilmasangiz, buni tekshirishingiz mumkin.
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// Bu URL bilan mos keluvchi base64 formati yordamida
	// kodlaydi/dekodlaydi.
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
