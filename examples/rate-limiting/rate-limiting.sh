# Dasturimizni ishga tushirsak, so'rovlarning birinchi
# to'plami xohlaganimizdek har ~200 millisekundda bir marta
# qayta ishlanayotganini ko'ramiz.
$ go run rate-limiting.go
request 1 2012-10-19 00:38:18.687438 +0000 UTC
request 2 2012-10-19 00:38:18.887471 +0000 UTC
request 3 2012-10-19 00:38:19.087238 +0000 UTC
request 4 2012-10-19 00:38:19.287338 +0000 UTC
request 5 2012-10-19 00:38:19.487331 +0000 UTC

# So'rovlarning ikkinchi to'plami uchun to'lqinli cheklash
# tufayli dastlabki 3 tasini darhol xizmat ko'rsatamiz,
# so'ngra qolgan 2 tasini har biri ~200ms kechikish bilan
# xizmat ko'rsatamiz.
request 1 2012-10-19 00:38:20.487578 +0000 UTC
request 2 2012-10-19 00:38:20.487645 +0000 UTC
request 3 2012-10-19 00:38:20.487676 +0000 UTC
request 4 2012-10-19 00:38:20.687483 +0000 UTC
request 5 2012-10-19 00:38:20.887542 +0000 UTC
