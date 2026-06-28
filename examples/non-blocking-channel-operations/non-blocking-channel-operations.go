// Kanallardagi oddiy yuborish va qabul qilishlar bloklovchidir.
// Biroq, biz `select` ni `default` bandi bilan ishlatib,
// _bloklamaydigan_ yuborish, qabul qilish va hatto
// bloklamaydigan ko'p yo'nalishli `select` larni amalga
// oshirishimiz mumkin.

package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// Mana bloklamaydigan qabul qilish. Agar `messages` da
	// qiymat mavjud bo'lsa, `select` o'sha qiymat bilan
	// `<-messages` `case` ini oladi. Agar bo'lmasa, u darhol
	// `default` case ini oladi.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// Bloklamaydigan yuborish ham xuddi shunday ishlaydi. Bu
	// yerda `msg` ni `messages` kanaliga yuborib bo'lmaydi,
	// chunki kanalda bufer yo'q va qabul qiluvchi ham yo'q.
	// Shuning uchun `default` case tanlanadi.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// Biz `default` bandi ustida bir nechta `case` dan
	// foydalanib, ko'p yo'nalishli bloklamaydigan select ni
	// amalga oshirishimiz mumkin. Bu yerda biz ham `messages`
	// ham `signals` da bloklamaydigan qabul qilishlarga harakat
	// qilamiz.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
