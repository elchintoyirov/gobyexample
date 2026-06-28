# Ushbu dasturni ishga tushirganimizda, avval bloklovchi
# chaqiruvning natijasini, so'ngra ikkita goroutinaning
# natijasini ko'ramiz. Goroutinalarning natijasi
# aralashib ketishi mumkin, chunki goroutinalar Go
# runtime tomonidan parallel ishga tushiriladi.
$ go run goroutines.go
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
done

# Keyingi misolda biz parallel Go dasturlarida
# goroutinalarni to'ldiruvchi vositani ko'rib chiqamiz:
# kanallar.
