# Dasturimizni ishga tushirish goroutinaga asoslangan holatni
# boshqarish misoli jami taxminan 80,000 ta amalni
# bajarishini ko'rsatadi.
$ go run stateful-goroutines.go
readOps: 71708
writeOps: 7177

# Bu aniq holatda goroutinaga asoslangan yondashuv mutexga
# asoslanganidan biroz murakkabroq bo'ldi. Biroq u ayrim
# holatlarda foydali bo'lishi mumkin, masalan, boshqa kanallar
# ishtirok etganda yoki bir nechta shunday mutexlarni boshqarish
# xatolarga olib kelishi mumkin bo'lganda. Siz qaysi yondashuv
# eng tabiiy his qilinsa, ayniqsa dasturingiz to'g'riligini
# tushunish nuqtai nazaridan, o'shanidan foydalanishingiz kerak.
