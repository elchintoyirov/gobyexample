// Go shablonga asoslangan layoutlar orqali vaqtni
// formatlash va tahlil qilishni qo'llab-quvvatlaydi.

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// Quyida vaqtni RFC3339 ga muvofiq, mos keladigan
	// layout konstantasidan foydalanib formatlashning
	// oddiy misoli keltirilgan.
	t := time.Now()
	p(t.Format(time.RFC3339))

	// Vaqtni tahlil qilish `Format` bilan bir xil layout
	// qiymatlaridan foydalanadi.
	t1, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	p(t1)

	// `Format` va `Parse` misolga asoslangan layoutlardan
	// foydalanadi. Odatda bu layoutlar uchun `time` paketidan
	// konstanta ishlatasiz, lekin o'zingiz xohlagan layoutlarni
	// ham berishingiz mumkin. Layoutlar berilgan vaqt/satrni
	// formatlash/tahlil qilish shablonini ko'rsatish uchun
	// `Mon Jan 2 15:04:05 MST 2006` etalon vaqtidan foydalanishi
	// kerak. Misol vaqti aynan ko'rsatilganidek bo'lishi shart:
	// 2006-yil, soat uchun 15, haftaning kuni uchun Monday va
	// hokazo.
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	p(t2)

	// Faqat raqamli ko'rinishlar uchun vaqt qiymatining
	// ajratib olingan tarkibiy qismlari bilan standart satr
	// formatlashdan ham foydalanishingiz mumkin.
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// Noto'g'ri kiritilgan ma'lumotda `Parse` tahlil qilish
	// muammosini tushuntiruvchi xato qaytaradi.
	_, err := time.Parse("Mon Jan _2 15:04:05 2006", "8:41PM")
	p(err)
}
