# Dasturni ishga tushirganimizda `"ping"` xabari kanalimiz
# orqali bir goroutinadan boshqasiga muvaffaqiyatli
# uzatiladi.
$ go run channels.go 
ping

# Standart holatda jo'natish va qabul qilish ham
# jo'natuvchi, ham qabul qiluvchi tayyor bo'lguncha
# bloklanadi. Ushbu xususiyat bizga dasturimiz oxirida
# boshqa hech qanday sinxronizatsiyadan foydalanmasdan
# `"ping"` xabarini kutish imkonini berdi.
