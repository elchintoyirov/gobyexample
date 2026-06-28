# Ishlayotgan dasturimiz 5 ta ishning turli ishchilar
# tomonidan bajarilayotganini ko'rsatadi. Dastur jami
# taxminan 5 soniyalik ish bajarishiga qaramay, faqat
# taxminan 2 soniya vaqt oladi, chunki 3 ta ishchi
# parallel ishlaydi.
$ time go run worker-pools.go 
worker 1 started  job 1
worker 2 started  job 2
worker 3 started  job 3
worker 1 finished job 1
worker 1 started  job 4
worker 2 finished job 2
worker 2 started  job 5
worker 3 finished job 3
worker 1 finished job 4
worker 2 finished job 5

real	0m2.358s
