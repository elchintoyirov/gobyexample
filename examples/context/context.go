// Oldingi misolda biz oddiy [HTTP server](http-server)
// sozlashni ko'rib chiqdik. HTTP serverlar bekor qilishni
// boshqarish uchun `context.Context` dan foydalanishni
// namoyish etishda foydalidir. `Context` deadline'larni,
// bekor qilish signallarini va boshqa so'rov doirasidagi qiymatlarni
// API chegaralari va goroutinalar bo'ylab olib yuradi.
package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	// Har bir so'rov uchun `context.Context` `net/http`
	// mexanizmi tomonidan yaratiladi va `Context()` metodi
	// orqali mavjud bo'ladi.
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	// Mijozga javob yuborishdan oldin bir necha soniya kutamiz.
	// Bu server bajarayotgan biror ishni simulyatsiya qilishi mumkin.
	// Ishlash davomida kontekstning `Done()` kanalini kuzatib boramiz:
	// u ishni bekor qilib, imkon qadar tezroq qaytishimiz kerakligi haqidagi
	// signalni beradi.
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		// Kontekstning `Err()` metodi `Done()` kanali nima uchun
		// yopilganini tushuntiruvchi xatoni qaytaradi.
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {

	// Avvalgidek, biz handlerimizni "/hello" yo'lida ro'yxatdan
	// o'tkazamiz va xizmat ko'rsatishni boshlaymiz.
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
