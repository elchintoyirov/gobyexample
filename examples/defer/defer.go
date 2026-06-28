// _Defer_ funksiya chaqiruvi dasturning bajarilishida
// keyinroq amalga oshirilishini ta'minlash uchun ishlatiladi,
// odatda tozalash maqsadlarida. `defer` ko'pincha boshqa tillarda
// masalan `ensure` va `finally` ishlatiladigan joylarda qo'llaniladi.

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Faraz qilaylik, biz fayl yaratmoqchimiz, unga yozmoqchimiz
// va ishimiz tugagach uni yopmoqchimiz. Buni `defer` bilan
// qanday qilishimiz mumkinligi quyida ko'rsatilgan.
func main() {

	// `createFile` bilan fayl obyektini olganimizdan so'ng
	// darhol o'sha faylni yopishni `closeFile` bilan defer qilamiz.
	// Bu o'rab turuvchi funksiya (`main`) oxirida, `writeFile`
	// tugaganidan keyin bajariladi.
	path := filepath.Join(os.TempDir(), "defer.txt")
	f := createFile(path)
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	// Faylni yopishda xatolarni tekshirish muhim, hatto
	// defer qilingan funksiyada ham.
	if err != nil {
		panic(err)
	}
}
