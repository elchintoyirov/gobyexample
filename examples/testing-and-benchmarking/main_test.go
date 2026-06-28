// Unit testlash printsipial Go dasturlarini yozishning muhim
// qismidir. `testing` paketi bizga unit testlar yozish uchun
// kerakli vositalarni taqdim etadi, `go test` buyrug'i esa
// testlarni ishga tushiradi.

// Namoyish maqsadida bu kod `main` paketida, lekin u istalgan
// paket bo'lishi mumkin. Test kodi odatda u test qiladigan kod
// bilan bir xil paketda joylashadi.
package main

import (
	"fmt"
	"testing"
)

// Biz butun sonlar minimumining bu oddiy implementatsiyasini
// test qilamiz. Odatda, biz test qilayotgan kod `intutils.go`
// kabi nomlangan manba faylda bo'lar edi, va u uchun test fayli
// esa `intutils_test.go` deb nomlangan bo'lardi.
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Test `Test` bilan boshlanadigan nomli funksiya yozish orqali
// yaratiladi.
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// `t.Error*` test muvaffaqiyatsizliklarini xabar qiladi,
		// lekin testni bajarishni davom ettiradi. `t.Fatal*` test
		// muvaffaqiyatsizliklarini xabar qiladi va testni darhol
		// to'xtatadi.
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// Test yozish takrorlanuvchan bo'lishi mumkin, shuning uchun
// *jadvalga asoslangan uslub*dan foydalanish odatiy holdir, unda
// test kirishlari va kutilgan chiqishlar jadvalda keltiriladi va
// bitta sikl ular bo'ylab yurib test mantiqini bajaradi.
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		// `t.Run` "subtest"larni, har bir jadval yozuvi uchun
		// bittadan, ishga tushirish imkonini beradi. Bular
		// `go test -v` bajarilganda alohida ko'rsatiladi.
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// Benchmark testlari odatda `_test.go` fayllarida bo'ladi va
// `Benchmark` bilan boshlanadigan nom bilan nomlanadi.
// Benchmark ishlashi uchun zarur, lekin o'lchanmasligi kerak
// bo'lgan har qanday kod bu sikldan oldin keladi.
func BenchmarkIntMin(b *testing.B) {
	for b.Loop() {
		// Benchmark ishga tushiruvchisi bitta iteratsiyaning
		// ishlash vaqtini oqilona baholash uchun bu sikl tanasini
		// avtomatik ravishda ko'p marta bajaradi.
		IntMin(1, 2)
	}
}
