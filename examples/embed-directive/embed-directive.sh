# Misolni ishga tushirish uchun ushbu buyruqlardan foydalaning.
# (Eslatma: go playground'dagi cheklov tufayli bu misolni
# faqat o'z lokal mashinangizda ishga tushirish mumkin.)
$ mkdir -p folder
$ echo "hello go" > folder/single_file.txt
$ echo "123" > folder/file1.hash
$ echo "456" > folder/file2.hash

$ go run embed-directive.go
hello go
hello go
123
456

