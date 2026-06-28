# Biz aniq 50,000 ta amal bo'lishini kutamiz. Agar atomik
# bo'lmagan butun sondan foydalanib, uni `ops++` bilan
# oshirganimizda edi, ehtimol har gal ishga tushirishda
# o'zgarib turadigan boshqacha son olardik, chunki
# goroutinalar bir-biriga xalal berardi. Bundan tashqari,
# `-race` bayrog'i bilan ishga tushirganimizda data race
# xatoliklariga duch kelardik.
$ go run atomic-counters.go
ops: 50000

# Keyingi misolda holatni boshqarishning yana bir vositasi
# bo'lgan mutexlarni ko'rib chiqamiz.
