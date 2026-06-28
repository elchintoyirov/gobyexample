// Biz ko'pincha Go kodini kelajakda biror nuqtada yoki
// biror oraliqda takror-takror bajarishni xohlaymiz. Go ning
// o'rnatilgan _taymer_ va _ticker_ imkoniyatlari bu
// vazifalarning ikkalasini ham osonlashtiradi. Avval
// taymerlarni, so'ngra [tickerlarni](tickers) ko'rib chiqamiz.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Taymerlar kelajakdagi bitta hodisani ifodalaydi. Siz
	// taymerga qancha kutmoqchi ekanligingizni aytasiz va u
	// o'sha vaqtda xabar beriladigan kanalni taqdim etadi.
	// Bu taymer 2 soniya kutadi.
	timer1 := time.NewTimer(2 * time.Second)

	// `<-timer1.C` taymerning `C` kanalida taymer
	// ishga tushganini bildiruvchi qiymat yuborilguncha
	// bloklaydi.
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// Agar shunchaki kutmoqchi bo'lsangiz, `time.Sleep`
	// dan foydalanishingiz mumkin edi. Taymer foydali
	// bo'lishining bir sababi shundaki, uni ishga
	// tushishidan oldin bekor qilishingiz mumkin.
	// Quyida shunga misol keltirilgan.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// `timer2` aslida to'xtatilganini ko'rsatish uchun,
	// agar u umuman ishga tushadigan bo'lsa, unga ishga
	// tushishga yetarli vaqt beramiz.
	time.Sleep(2 * time.Second)
}
