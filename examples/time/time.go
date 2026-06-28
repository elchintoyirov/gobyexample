// Go vaqtlar va davomiyliklar uchun keng qamrovli
// qo'llab-quvvatlashni taklif qiladi; quyida bir nechta
// misol keltirilgan.

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// Joriy vaqtni olishdan boshlaymiz.
	now := time.Now()
	p(now)

	// Yil, oy, kun va hokazolarni berib `time` struct
	// yaratishingiz mumkin. Vaqtlar har doim `Location`,
	// ya'ni vaqt mintaqasi bilan bog'liq bo'ladi.
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// Kutilganidek, vaqt qiymatining turli tarkibiy
	// qismlarini ajratib olishingiz mumkin.
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// Dushanba-Yakshanba oralig'idagi `Weekday` ham mavjud.
	p(then.Weekday())

	// Bu metodlar ikki vaqtni taqqoslab, birinchisi
	// ikkinchisidan oldin, keyin yoki ayni o'sha vaqtda
	// yuz berishini mos ravishda tekshiradi.
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// `Sub` metodi ikki vaqt orasidagi oraliqni
	// ifodalovchi `Duration` qaytaradi.
	diff := now.Sub(then)
	p(diff)

	// Davomiylik uzunligini turli birliklarda
	// hisoblashimiz mumkin.
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// Vaqtni berilgan davomiylikka oldinga surish uchun
	// `Add` dan, yoki davomiylikka orqaga qaytarish uchun
	// `-` bilan ishlatishingiz mumkin.
	p(then.Add(diff))
	p(then.Add(-diff))
}
