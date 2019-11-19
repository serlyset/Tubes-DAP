package main

import (
	"fmt"
)

type Movie struct {
	Judul string
	Durasi string
	Harga int
	Views int
}
film [3] Movie
Studio [5] Schedule
//Seat [5][5] bool
Schedule [5] film

type Snack struct {
	Jenis string
	Harga int
	Ukuran string
}

func CetakJadwal (a int) {
	fmt.Println("")
}

func main() {
	var (
		stu Studio
		sch Schedule
		fil film
	)
	fil[0].Judul = "Joker"
	fil[0].Durasi = "122 menit"
	fil[0].Harga = 20000
	fil[0].Views = 0

	fil[1].Judul = "Toy Story 4"
	fil[1].Durasi = "100 menit"
	fil[1].Harga = 20000
	fil[1].Views = 0

	fil[2].Judul = "Avengers : Endgame"
	fil[2].Durasi = "182 menit"
	fil[2].Harga = 20000
	fil[2].Views = 0

	for i:=0; i<5; i++ {
		Studio[i].s
	}

	fmt.Println("                       BIOSKOP BOJONGSOANG")
	fmt.Println("                 Jalan Terusan Buah Batu, Bandung")
	fmt.Println("----------------------------------------------------------------------")
	fmt.Printf("\nMENU : \n")
	fmt.Println("1. Lihat Jadwal")
	fmt.Println("2. Beli tiket")
	fmt.Println("3. Beli snack")
	fmt.Println("0. Keluar")

	var choice = 99
	fmt.Print("Masukkan Pilihan Anda : ")
	fmt.Scanln(&choice)
	switch choice {
	case 1:

	case 2:

	case 3:

	case 0:
		fmt.Println("----------Terima kasih telah mengunjungi Bioskop Bojongsoang----------")
		fmt.Println("                    Kami tunggu kedatangannya kembali")
	}
}
