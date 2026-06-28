# Faylga yozish kodini ishga tushirib ko'ring.
$ go run writing-files.go 
wrote 5 bytes
wrote 7 bytes
wrote 9 bytes

# Keyin yozilgan fayllar tarkibini tekshiring.
$ cat /tmp/dat1
hello
go
$ cat /tmp/dat2
some
writes
buffered

# Keyingi qadamda biz hozir ko'rgan fayl I/O g'oyalarining
# ba'zilarini `stdin` va `stdout` oqimlariga qo'llashni
# ko'rib chiqamiz.
