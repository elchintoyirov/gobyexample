# Namuna chiqish; chiqarilgan sana va vaqt misol qachon
# ishga tushirilganiga bog'liq bo'ladi.
$ go run logging.go
2023/08/22 10:45:16 standard logger
2023/08/22 10:45:16.904141 with micro
2023/08/22 10:45:16 logging.go:40: with file/line
my:2023/08/22 10:45:16 from mylog
ohmy:2023/08/22 10:45:16 from mylog
from buflog:buf:2023/08/22 10:45:16 hello

# Bular veb-saytda taqdimot aniqligi uchun bir nechta
# qatorga bo'lingan; aslida ular bitta qatorda chiqariladi.
{"time":"2023-08-22T10:45:16.904166391-07:00",
 "level":"INFO","msg":"hi there"}
{"time":"2023-08-22T10:45:16.904178985-07:00",
	"level":"INFO","msg":"hello again",
	"key":"val","age":25}
