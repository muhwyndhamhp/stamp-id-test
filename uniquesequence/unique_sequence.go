package uniquesequence

import (
	"fmt"
)

func UniquePrintSequence() {
	var result Stack
	for i := 1; i <= 100; i++ {
		if isPrime(i) {
			continue
		}

		numeric := ""

		if i%3 == 0 {
			numeric += "Foo"
		}

		if i%5 == 0 {
			numeric += "Bar"
		}

		if numeric == "" {
			numeric = fmt.Sprintf("%d", i)
		}

		result.Push(numeric)
	}

	length := len(result)

	for i := 0; i < length; i++ {
		if i+1 < length {
			fmt.Printf("%s, ", result.Pop())
		} else {
			fmt.Printf("%s", result.Pop())
		}

	}
}

// Logic:
// Bilangan prima adalah bilangan yang memenuhi kriteria:
// - Bilangan > 1
// - Bilangan hanya bisa habis dibagi oleh 1 dan dirinya sendiri (misal 2 hanya bisa dibagi angka 1 dan 2)
// Oleh karena itu, langkah pengecekan bisa dimulai dengan memastikan angkanya > 1
// Dilanjutkan dengan membuat pembagi yang nilai perkalian pembaginya < nilai angkanya
// (misal 3 * 3  = 9; 9 lebih dari 4 maka tidak valid)
// kemudian dicek apakah bilangan pembagi tersebut habis membagi bilangan yang ingin kita cek prima atau bukan
// Jika ada bilangan yang habis membagi bilangan tersebut, kita anggap bukan bilangan prima
func isPrime(num int) bool {

	if num <= 1 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

type Stack []string

func (s *Stack) Push(v string) {
	*s = append(*s, v)
}

func (s *Stack) Pop() string {
	ret := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]

	return ret
}
