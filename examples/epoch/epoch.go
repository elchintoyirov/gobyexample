// Dasturlardagi keng tarqalgan talab [Unix
// epoch](https://en.wikipedia.org/wiki/Unix_time)dan beri
// o'tgan soniyalar, millisoniyalar yoki nanosoniyalar sonini
// olishdir. Buni Go'da qanday qilish quyida ko'rsatilgan.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Unix epoch'dan beri o'tgan vaqtni mos ravishda soniya,
	// millisoniya yoki nanosoniyalarda olish uchun `time.Now` ni
	// `Unix`, `UnixMilli` yoki `UnixNano` bilan ishlating.
	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	// Shuningdek, epoch'dan beri o'tgan butun sonli soniya yoki
	// nanosoniyalarni mos keluvchi `time` ga aylantirishingiz mumkin.
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}
