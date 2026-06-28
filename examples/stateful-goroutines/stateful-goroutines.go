// Oldingi misolda biz bir nechta goroutina o'rtasidagi umumiy
// holatdan foydalanishni sinxronlash uchun [mutexlar](mutexes)
// yordamida oshkora qulflashdan foydalandik. Yana bir variant -
// xuddi shu natijaga erishish uchun goroutinalar va kanallarning
// o'rnatilgan sinxronlash imkoniyatlaridan foydalanishdir. Bu
// kanalga asoslangan yondashuv Go'ning xotirani aloqa orqali
// bo'lishish va har bir ma'lumot bo'lagini aniq bitta goroutina
// egaligida saqlash g'oyalariga mos keladi.

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// Bu misolda bizning holatimizning egasi bitta goroutina bo'ladi.
// Bu ma'lumot bir vaqtning o'zida foydalanish bilan hech qachon
// buzilmasligini kafolatlaydi. Bu holatni o'qish yoki yozish
// uchun boshqa goroutinalar egasi bo'lgan goroutinaga xabarlar
// yuboradi va mos javoblarni qabul qiladi. Bu `readOp` va
// `writeOp` `struct`lari o'sha so'rovlarni hamda egasi bo'lgan
// goroutina javob berishi uchun yo'lni qamrab oladi.
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	// Avvalgidek biz qancha amal bajarayotganimizni sanaymiz.
	var readOps uint64
	var writeOps uint64

	// `reads` va `writes` kanallari boshqa goroutinalar tomonidan
	// mos ravishda o'qish va yozish so'rovlarini yuborish uchun
	// ishlatiladi.
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// Mana `state`ning egasi bo'lgan goroutina, u oldingi
	// misoldagi kabi map, lekin endi holatli goroutinaga xos.
	// Bu goroutina `reads` va `writes` kanallarida takror-takror
	// select qiladi va so'rovlar kelganda ularga javob beradi.
	// Javob avval so'ralgan amalni bajarish, so'ngra muvaffaqiyatni
	// bildirish uchun `resp` javob kanaliga qiymat yuborish bilan
	// amalga oshiriladi (`reads` holatida kerakli qiymat ham
	// yuboriladi).
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// Bu `reads` kanali orqali holat egasi bo'lgan goroutinaga
	// o'qish so'rovlarini yuborish uchun 100 ta goroutina ishga
	// tushiradi. Har bir o'qish `readOp` tuzishni, uni `reads`
	// kanali orqali yuborishni, so'ngra berilgan `resp` kanali
	// orqali natijani qabul qilishni talab qiladi.
	for range 100 {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Biz shunga o'xshash yondashuvdan foydalanib 10 ta yozishni
	// ham ishga tushiramiz.
	for range 10 {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Goroutinalar bir soniya ishlashiga ruxsat beramiz.
	time.Sleep(time.Second)

	// Nihoyat, amal sonlarini olib chiqamiz va chop etamiz.
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
