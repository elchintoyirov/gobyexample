// `net/http` paketidan foydalanib asosiy HTTP server
// yozish oson.
package main

import (
	"fmt"
	"net/http"
)

// `net/http` serverlarida asosiy tushuncha bu
// *ishlovchilar*dir. Ishlovchi - bu `http.Handler`
// interfeysini amalga oshiruvchi obyekt. Ishlovchi
// yozishning keng tarqalgan usuli - mos imzoga ega
// funksiyalarda `http.HandlerFunc` adapteridan
// foydalanishdir.
func hello(w http.ResponseWriter, req *http.Request) {

	// Ishlovchi sifatida xizmat qiluvchi funksiyalar
	// argument sifatida `http.ResponseWriter` va
	// `http.Request` ni qabul qiladi. Response writer HTTP
	// javobini to'ldirish uchun ishlatiladi. Bu yerda bizning
	// oddiy javobimiz shunchaki "hello\n".
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	// Bu ishlovchi barcha HTTP so'rov sarlavhalarini o'qib,
	// ularni javob tanasiga aks ettirib, biroz murakkabroq
	// ish bajaradi.
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	// Biz ishlovchilarimizni server marshrutlarida
	// `http.HandleFunc` qulay funksiyasidan foydalanib
	// ro'yxatdan o'tkazamiz. U `net/http` paketida *standart
	// marshrutlovchini* o'rnatadi va argument sifatida
	// funksiyani qabul qiladi.
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// Nihoyat, biz `ListenAndServe` ni port va ishlovchi
	// bilan chaqiramiz. `nil` unga biz hozirgina o'rnatgan
	// standart marshrutlovchidan foydalanishni aytadi.
	http.ListenAndServe(":8090", nil)
}
