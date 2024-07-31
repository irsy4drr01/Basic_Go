package task

import "fmt"

func Task_3() {
	var sisi float64

	// Mengambil input dari pengguna untuk sisi
	for {
		fmt.Print("\nMasukkan panjang sisi: ")
		_, err := fmt.Scan(&sisi)
		if err != nil {
			fmt.Println("Error:\nInput bukan angka. Silakan masukkan angka!")

			var clearInput string
			fmt.Scanln(&clearInput)
		} else {
			break
		}
	}
	fmt.Println()

	// Membuat objek Persegi dan Kubus dengan input dari pengguna
	persegi := Persegi{sisi: sisi}
	kubus := Kubus{sisi: sisi}

	// Menyimpan objek ke dalam variabel bertipe interface
	// Menetapkan objek ke dalam interface
	var h2d hitung2d = &persegi
	var h3d hitung3d = &kubus
	var h hitung = &kubus

	// Menampilkan hasil perhitungan untuk Persegi
	fmt.Printf("Persegi:\n")
	fmt.Printf("Luas: %.2f\n", h2d.luas())
	fmt.Printf("Keliling: %.2f\n", h2d.keliling())

	// Menampilkan hasil perhitungan untuk Kubus
	fmt.Printf("\nKubus:\n")
	fmt.Printf("Luas: %.2f\n", h.luas())
	fmt.Printf("Keliling: %.2f\n", h.keliling())
	fmt.Printf("Volume: %.2f\n", h3d.volume())

	fmt.Println()
}

// Mendefinisikan interface hitung2d
type hitung2d interface {
	luas() float64
	keliling() float64
}

// Mendefinisikan interface hitung3d
type hitung3d interface {
	volume() float64
}

// Mendefinisikan interface hitung yang menggabungkan hitung2d dan hitung3d
type hitung interface {
	hitung2d
	hitung3d
}

// Struktur Persegi yang mengimplementasikan interface hitung2d
type Persegi struct {
	sisi float64
}

// Implementasi metode luas untuk Persegi
func (p *Persegi) luas() float64 {
	return p.sisi * p.sisi
}

// Implementasi metode keliling untuk Persegi
func (p *Persegi) keliling() float64 {
	return 4 * p.sisi
}

// Struktur Kubus yang mengimplementasikan interface hitung2d dan hitung3d
type Kubus struct {
	sisi float64
}

// Implementasi metode luas untuk Kubus (untuk hitung2d)
func (k *Kubus) luas() float64 {
	return 6 * k.sisi * k.sisi
}

// Implementasi metode keliling untuk Kubus (untuk hitung2d)
func (k *Kubus) keliling() float64 {
	return 12 * k.sisi
}

// Implementasi metode volume untuk Kubus (untuk hitung3d)
func (k *Kubus) volume() float64 {
	return k.sisi * k.sisi * k.sisi
}
