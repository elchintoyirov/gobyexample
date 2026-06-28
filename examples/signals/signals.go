// Ba'zan Go dasturlarimiz [Unix signallarini](https://en.wikipedia.org/wiki/Unix_signal)
// aqlli tarzda qayta ishlashini xohlaymiz. Masalan, server
// `SIGTERM` qabul qilganda muloyimlik bilan o'chishini, yoki
// buyruq qatori vositasi `SIGINT` qabul qilsa kirishni qayta
// ishlashni to'xtatishini xohlashimiz mumkin. Mana kontekstlar
// yordamida signallarni qayta ishlashning zamonaviy usuli.

package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
)

func main() {
	// `signal.NotifyContext` sanab o'tilgan signallardan biri
	// kelganda bekor qilinadigan kontekstni qaytaradi.
	ctx, stop := signal.NotifyContext(
		context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Dastur sozlangan signallardan biri qabul qilinmaguncha
	// shu yerda kutadi.
	fmt.Println("awaiting signal")
	<-ctx.Done()

	// `context.Cause` kontekst nima uchun bekor qilinganini
	// xabar qiladi. Signal tomonidan qo'zg'atilgan bekor
	// qilish uchun bu signal qiymatini o'z ichiga oladi.
	fmt.Println()
	fmt.Println(context.Cause(ctx))
	fmt.Println("exiting")
}
