// `net` paketi TCP soket serverlarini oson qurish uchun bizga
// kerakli vositalarni taqdim etadi.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	// `net.Listen` serverni berilgan tarmoq (TCP) va manzilda
	// (barcha interfeyslardagi 8090 port) ishga tushiradi.
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal("Error listening:", err)
	}

	// Dastur chiqib ketganda portni bo'shatish uchun listenerni
	// yoping.
	defer listener.Close()

	// Yangi mijoz ulanishlarini qabul qilish uchun cheksiz sikl.
	for {
		// Ulanishni kuting.
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting conn:", err)
			continue
		}

		// Bu yerda biz ulanishni boshqarish uchun goroutinadan
		// foydalanamiz, shunda asosiy sikl yana ko'proq
		// ulanishlarni qabul qilishni davom ettira oladi.
		go handleConnection(conn)
	}
}

// `handleConnection` bitta mijoz ulanishini boshqaradi, mijozdan
// bir qator matnni o'qiydi va javob qaytaradi.
func handleConnection(conn net.Conn) {
	// Mijoz bilan o'zaro aloqani tugatganimizda ulanishni yopish
	// resurslarni bo'shatadi.
	defer conn.Close()

	// Mijozdan bir qator ma'lumotni (yangi qator bilan
	// tugaydigan) o'qish uchun `bufio.NewReader` dan foydalaning.
	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Read error: %v", err)
		return
	}

	// Ikki tomonlama aloqani namoyish etib, mijozga javob
	// yaratib qaytarib yuboring.
	ackMsg := strings.ToUpper(strings.TrimSpace(message))
	response := fmt.Sprintf("ACK: %s\n", ackMsg)
	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Printf("Server write error: %v", err)
	}
}
