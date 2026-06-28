// [_SHA256 hashlari_](https://en.wikipedia.org/wiki/SHA-2)
// binary yoki matnli bloblar uchun qisqa identifikatorlarni
// hisoblashda tez-tez ishlatiladi. Masalan, TLS/SSL
// sertifikatlari sertifikat imzosini hisoblash uchun SHA256dan
// foydalanadi. Mana Go'da SHA256 hashlarini qanday hisoblash.

package main

// Go bir nechta hash funksiyalarini turli `crypto/*`
// paketlarida amalga oshiradi.
import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 this string"

	// Bu yerda yangi hash bilan boshlaymiz.
	h := sha256.New()

	// `Write` baytlarni kutadi. Agar sizda `s` satri bo'lsa,
	// uni baytlarga aylantirish uchun `[]byte(s)`dan foydalaning.
	h.Write([]byte(s))

	// Bu yakuniy hash natijasini bayt slice sifatida oladi.
	// `Sum`ga berilgan argument mavjud bayt slice'iga
	// qo'shish uchun ishlatilishi mumkin: odatda u kerak emas.
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
