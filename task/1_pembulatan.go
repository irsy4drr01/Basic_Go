package task

import (
	"fmt"
	"math"
)

func Task_1() {
	var N float64

	for {
		fmt.Print("\nInput value N: ")
		_, err := fmt.Scan(&N)
		if err != nil {
			fmt.Println("Error:\nInput bukan angka. Silahkan masukkan angka!")

			clearInput := ""
			fmt.Println(clearInput)
		} else {
			break
		}
	}
	result := pembulatan(N)
	fmt.Printf("Hasil pembulatan %f adalah %.2f", N, result)
	fmt.Println()
}

func pembulatan(N float64) float64 {
	return math.Round(N*10) / 10
}
