// Go standart kutubxonasi Go dasturlaridan loglarni
// chiqarish uchun sodda vositalarni taqdim etadi: erkin
// shakldagi chiqish uchun [log](https://pkg.go.dev/log)
// paketi va tuzilmali chiqish uchun
// [log/slog](https://pkg.go.dev/log/slog) paketi.
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"log/slog"
)

func main() {

	// `log` paketidan `Println` kabi funksiyalarni oddiy
	// chaqirish _standart_ loggerdan foydalanadi, u allaqachon
	// `os.Stderr` ga oqilona log chiqarish uchun oldindan
	// sozlangan. `Fatal*` yoki `Panic*` kabi qo'shimcha
	// metodlar logdan keyin dasturdan chiqadi.
	log.Println("standard logger")

	// Loggerlar chiqish formatini belgilash uchun _bayroqlar_
	// bilan sozlanishi mumkin. Standart holatda, standart
	// loggerda `log.Ldate` va `log.Ltime` bayroqlari
	// o'rnatilgan va ular `log.LstdFlags` da to'plangan.
	// Masalan, biz uning bayroqlarini vaqtni mikrosekund
	// aniqligida chiqarish uchun o'zgartirishimiz mumkin.
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	// U, shuningdek, `log` funksiyasi chaqirilgan fayl nomi
	// va qatorni chiqarishni qo'llab-quvvatlaydi.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	// Maxsus logger yaratish va uni uzatish foydali bo'lishi
	// mumkin. Yangi logger yaratganda, biz uning chiqishini
	// boshqa loggerlardan ajratish uchun _prefiks_
	// o'rnatishimiz mumkin.
	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")

	// Biz mavjud loggerlarda (shu jumladan standart loggerda)
	// prefiksni `SetPrefix` metodi bilan o'rnatishimiz mumkin.
	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	// Loggerlar maxsus chiqish maqsadlariga ega bo'lishi
	// mumkin; istalgan `io.Writer` ishlaydi.
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)

	// Bu chaqiruv log chiqishini `buf` ga yozadi.
	buflog.Println("hello")

	// Bu aslida uni standart chiqishda ko'rsatadi.
	fmt.Print("from buflog:", buf.String())

	// `slog` paketi _tuzilmali_ log chiqishini taqdim etadi.
	// Masalan, JSON formatida loglash sodda.
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")

	// Xabardan tashqari, `slog` chiqishi ixtiyoriy sondagi
	// key=value juftliklarini o'z ichiga olishi mumkin.
	myslog.Info("hello again", "key", "val", "age", 25)
}
