# Satr standart va URL base64 enkoderlari bilan biroz
# farqli qiymatlarga kodlanadi (oxiridagi `+` va `-`), lekin
# ularning ikkalasi ham kerakli tarzda asl satrga
# dekodlanadi.
$ go run base64-encoding.go
YWJjMTIzIT8kKiYoKSctPUB+
abc123!?$*&()'-=@~

YWJjMTIzIT8kKiYoKSctPUB-
abc123!?$*&()'-=@~
