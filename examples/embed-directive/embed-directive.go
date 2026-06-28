// `//go:embed` bu [kompilyator
// direktivasi](https://pkg.go.dev/cmd/compile#hdr-Compiler_Directives) bo'lib,
// dasturlarga build vaqtida Go binar fayliga ixtiyoriy fayllar va papkalarni
// kiritish imkonini beradi. Embed direktivasi haqida ko'proq
// [bu yerda](https://pkg.go.dev/embed) o'qing.
package main

// `embed` paketini import qilamiz; agar bu paketdan hech qanday
// eksport qilingan identifikatordan foydalanmasangiz, `_ "embed"` bilan
// bo'sh import qilishingiz mumkin.
import (
	"embed"
)

// `embed` direktivalari Go manba faylini o'z ichiga olgan katalogga
// nisbatan yo'llarni qabul qiladi. Bu direktiva faylning tarkibini
// undan keyin darhol keladigan `string` o'zgaruvchisiga joylashtiradi.
//
//go:embed folder/single_file.txt
var fileString string

// Yoki faylning tarkibini `[]byte` ga joylashtiramiz.
//
//go:embed folder/single_file.txt
var fileByte []byte

// Biz wildcard'lar bilan bir nechta fayl yoki hatto papkalarni ham
// joylashtirishimiz mumkin. Bu oddiy virtual fayl tizimini amalga oshiradigan
// [embed.FS tipi](https://pkg.go.dev/embed#FS) o'zgaruvchisidan foydalanadi.
//
//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {

	// `single_file.txt` tarkibini chop etamiz.
	print(fileString)
	print(string(fileByte))

	// Joylashtirilgan papkadan bir nechta faylni olamiz.
	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}
