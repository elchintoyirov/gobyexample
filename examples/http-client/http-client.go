// Go standart kutubxonasi `net/http` paketida HTTP
// mijozlari va serverlari uchun ajoyib qo'llab-quvvatlash
// bilan birga keladi. Ushbu misolda biz undan oddiy HTTP
// so'rovlarini yuborish uchun foydalanamiz.
package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	// Serverga HTTP GET so'rovini yuboring. `http.Get`
	// `http.Client` obyektini yaratish va uning `Get`
	// metodini chaqirish atrofidagi qulay qisqartmadir; u
	// foydali standart sozlamalarga ega `http.DefaultClient`
	// obyektidan foydalanadi.
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// HTTP javob statusini chop eting.
	fmt.Println("Response status:", resp.Status)

	// Javob tanasining dastlabki 5 qatorini chop eting.
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
