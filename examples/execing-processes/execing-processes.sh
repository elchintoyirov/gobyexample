# Dasturimizni ishga tushirganimizda u `ls` bilan almashtiriladi.
$ go run execing-processes.go
total 16
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 execing-processes.go

# E'tibor bering, Go klassik Unix `fork` funksiyasini
# taklif qilmaydi. Biroq odatda bu muammo emas, chunki
# goroutinalarni ishga tushirish, jarayonlarni yaratish
# va jarayonlarni exec qilish `fork` ning ko'pchilik
# foydalanish holatlarini qamrab oladi.
