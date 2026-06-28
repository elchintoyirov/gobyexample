// _Kanallar_ — bu parallel goroutinalarni bir-biriga
// bog'laydigan quvurlardir. Bir goroutinadan kanallarga
// qiymat jo'natib, o'sha qiymatlarni boshqa goroutinada
// qabul qilishingiz mumkin.

package main

import "fmt"

func main() {

	// `make(chan val-type)` bilan yangi kanal yarating.
	// Kanallar o'zlari uzatadigan qiymatlar tipiga ega
	// bo'ladi.
	messages := make(chan string)

	// `channel <-` sintaksisi yordamida kanalga qiymat
	// _jo'nating_. Bu yerda biz yuqorida yaratgan `messages`
	// kanaliga yangi goroutinadan `"ping"`ni jo'natamiz.
	go func() { messages <- "ping" }()

	// `<-channel` sintaksisi kanaldan qiymatni _qabul
	// qiladi_. Bu yerda biz yuqorida jo'natgan `"ping"`
	// xabarini qabul qilib, uni chop etamiz.
	msg := <-messages
	fmt.Println(msg)
}
