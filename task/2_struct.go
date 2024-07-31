package task

import (
	"fmt"
	"math"
)

func Task_2() {
	var limit int
	for {
		fmt.Print("\nMasukkan limit deret bilangan: ")
		_, err := fmt.Scan(&limit)
		if err != nil {
			fmt.Println("Error:\nInput bukan angka. Silakan masukkan angka!")

			clearInput := ""
			fmt.Println(clearInput)
		} else {
			break
		}
	}

	deret := DeretBilangan{limit}
	deret.prima()
	deret.ganjil()
	deret.genap()
	deret.fibonacci()
	fmt.Println()
}

// Struct untuk menyimpan limit deret bilangan
type DeretBilangan struct {
	limit int
}

// function untuk memeriksa apakah suatu bilangan adalah bilangan prima
func isPrima(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Method untuk mencetak deret bilangan prima
func (db *DeretBilangan) prima() {
	fmt.Printf("Bilangan Prima hingga %d:\n", db.limit)
	for i := 2; i <= db.limit; i++ {
		if isPrima(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}

// Method untuk mencetak deret bilangan ganjil
func (db *DeretBilangan) ganjil() {
	fmt.Printf("Bilangan Ganjil hingga %d:\n", db.limit)
	for i := 1; i <= db.limit; i += 2 {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

// Method untuk mencetak deret bilangan genap
func (db *DeretBilangan) genap() {
	fmt.Printf("Bilangan Genap hingga %d:\n", db.limit)
	for i := 2; i <= db.limit; i += 2 {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

// Method untuk mencetak deret bilangan Fibonacci
func (db *DeretBilangan) fibonacci() {
	fmt.Printf("Bilangan Fibonacci hingga %d:\n", db.limit)
	a, b := 0, 1
	for a <= db.limit {
		fmt.Printf("%d ", a)
		a, b = b, a+b
	}
	fmt.Println()
}
