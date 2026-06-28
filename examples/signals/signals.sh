# Bu dasturni ishga tushirsak, u signal kutib bloklanadi.
# `ctrl-C` (terminal uni `^C` sifatida ko'rsatadi) ni terib
# `SIGINT` signalini yubora olamiz, bu dasturni bekor qilish
# sababini chop etishga va keyin chiqishga majbur qiladi.
$ go run signals.go
awaiting signal
^C
interrupt signal received
exiting
