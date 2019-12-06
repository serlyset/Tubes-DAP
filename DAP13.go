package main

import (
	"fmt"
)

type Seat [5][5]int
type Movie struct {
	Judul  string
	Durasi string
	Harga  int
	Views  int
	Kursi  Seat
}
type SchContent [3]Movie
type film struct {
	Judul  string
	Durasi string
	Views  int
}

var Popular [3]film

var (
	sch1, sch2, sch3 SchContent
)

func SortPopular(sch1 *SchContent, sch2 *SchContent, sch3 *SchContent, Popular *[3]film) {
	//kalo ada view yang sama, slot awal dianggap lebih populer
	var v0, v1, v2 int
	v0 = sch1[0].Views + sch2[2].Views + sch3[1].Views
	v1 = sch1[1].Views + sch2[0].Views + sch3[2].Views
	v2 = sch1[2].Views + sch2[1].Views + sch3[0].Views

	if v0 >= v1 {
		if v1 >= v2 { //descending v0 v1 v2
			Popular[0].Judul = sch1[0].Judul
			Popular[0].Durasi = sch1[0].Durasi
			Popular[0].Views = v0

			Popular[1].Judul = sch1[1].Judul
			Popular[1].Durasi = sch1[1].Durasi
			Popular[1].Views = v1

			Popular[2].Judul = sch1[2].Judul
			Popular[2].Durasi = sch1[2].Durasi
			Popular[2].Views = v2
		} else if v0 >= v2 { //descending v0 v2 v1
			Popular[0].Judul = sch1[0].Judul
			Popular[0].Durasi = sch1[0].Durasi
			Popular[0].Views = v0

			Popular[1].Judul = sch1[2].Judul
			Popular[1].Durasi = sch1[2].Durasi
			Popular[1].Views = v2

			Popular[2].Judul = sch1[1].Judul
			Popular[2].Durasi = sch1[1].Durasi
			Popular[2].Views = v1
		} else { //descending v2 v0 v1
			Popular[0].Judul = sch1[2].Judul
			Popular[0].Durasi = sch1[2].Durasi
			Popular[0].Views = v2

			Popular[1].Judul = sch1[0].Judul
			Popular[1].Durasi = sch1[0].Durasi
			Popular[1].Views = v0

			Popular[2].Judul = sch1[1].Judul
			Popular[2].Durasi = sch1[1].Durasi
			Popular[2].Views = v1
		}
	} else {
		if v0 >= v2 { //descending v1 v0 v2
			Popular[0].Judul = sch1[1].Judul
			Popular[0].Durasi = sch1[1].Durasi
			Popular[0].Views = v1

			Popular[1].Judul = sch1[0].Judul
			Popular[1].Durasi = sch1[0].Durasi
			Popular[1].Views = v0

			Popular[2].Judul = sch1[2].Judul
			Popular[2].Durasi = sch1[2].Durasi
			Popular[2].Views = v2
		} else if v1 >= v2 { //descending v1 v2 v0
			Popular[0].Judul = sch1[1].Judul
			Popular[0].Durasi = sch1[1].Durasi
			Popular[0].Views = v1

			Popular[1].Judul = sch1[2].Judul
			Popular[1].Durasi = sch1[2].Durasi
			Popular[1].Views = v2

			Popular[2].Judul = sch1[0].Judul
			Popular[2].Durasi = sch1[0].Durasi
			Popular[2].Views = v0
		} else { // descending v2 v1 v0
			Popular[0].Judul = sch1[2].Judul
			Popular[0].Durasi = sch1[2].Durasi
			Popular[0].Views = v2

			Popular[1].Judul = sch1[1].Judul
			Popular[1].Durasi = sch1[1].Durasi
			Popular[1].Views = v1

			Popular[2].Judul = sch1[0].Judul
			Popular[2].Durasi = sch1[0].Durasi
			Popular[2].Views = v0
		}
	}
}

func PrintPopular(Popular [3]film) {
	fmt.Println("\nFilm paling populer dengan penjualan tiket terbanyak adalah:")
	fmt.Print("[1] ", Popular[0].Judul, "\n")
	fmt.Print("    Berdurasi ", Popular[0].Durasi, "\n")
	fmt.Print("    Tiket terjual: ", Popular[0].Views, "\n")

	fmt.Print("[2] ", Popular[1].Judul, "\n")
	fmt.Print("    Berdurasi ", Popular[1].Durasi, "\n")
	fmt.Print("    Tiket terjual: ", Popular[1].Views, "\n")

	fmt.Print("[3] ", Popular[2].Judul, "\n")
	fmt.Print("    Berdurasi ", Popular[2].Durasi, "\n")
	fmt.Print("    Tiket terjual: ", Popular[2].Views, "\n")
}

func Report() {
	fmt.Println()
	fmt.Println("Hasil penjualan tiket hari ini:")

	fmt.Print("[1] ", sch1[0].Judul, "\n")
	fmt.Print("    ", "Tiket terjual pada jadwal 10.00-13.00: ", sch1[0].Views, "\n")
	fmt.Print("    ", "Tiket terjual pada jadwal 14.00-17.00: ", sch2[2].Views, "\n")
	fmt.Print("    ", "Tiket terjual pada jadwal 18.00-21.00: ", sch3[1].Views, "\n")

	fmt.Print("[2] ", sch2[0].Judul, "\n")
	fmt.Print("    ", "Tiket terjual pada jadwal 10.00-13.00: ", sch1[1].Views, "\n")
	fmt.Print("    ", "Tiket terjual pada jadwal 14.00-17.00: ", sch2[0].Views, "\n")
	fmt.Print("    ", "Tiket terjual pada jadwal 18.00-21.00: ", sch3[2].Views, "\n")

	fmt.Print("[3] ", sch3[0].Judul, "\n")
	fmt.Print("    ", "Tiket terjual pada jadwal 10.00-13.00: ", sch1[2].Views, "\n")
	fmt.Print("    ", "Tiket terjual pada jadwal 14.00-17.00: ", sch2[1].Views, "\n")
	fmt.Print("    ", "Tiket terjual pada jadwal 18.00-21.00: ", sch3[0].Views, "\n")

	SortPopular(&sch1, &sch2, &sch3, &Popular)
	PrintPopular(Popular)

	MenuAdmin()
}

func ResetSeat(ar *Seat) { //done
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			ar[i][j] = 0
		}
	}
	ar[0][0] = 1
	ar[0][4] = 1
	ar[4][0] = 1
	ar[4][4] = 1
	ar[0][1] = 1
	ar[4][1] = 1
}

func ResetAllSeat(sch1 *SchContent, sch2 *SchContent, sch3 *SchContent) { //done
	for i := 0; i < 3; i++ {
		ResetSeat(&sch1[i].Kursi)
		ResetSeat(&sch2[i].Kursi)
		ResetSeat(&sch3[i].Kursi)
	}
	fmt.Println("Data kursi sudah berhasil direset. Mengalihkan ke menu admin...\n")
	MenuAdmin()
}

func FindEmpty(a Movie) (int, int) { //return index kursi
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if a.Kursi[i][j] == 0 {
				return i, j
			}
		}
	}
	return 99, 99
}

func SubCetakJadwal(p int, q string) { //done
	if q == "" {
		q = "-"
		p = 1
	}
	var a int
	a = 25 - p
	if p%2 == 1 {
		for i := 1; i <= a/2; i++ {
			fmt.Print(" ")
		}
		fmt.Print(q)
		for i := 1; i <= a/2; i++ {
			fmt.Print(" ")
		}
	} else {
		var b, c int
		b = a / 2
		c = a - b
		for i := 1; i <= b; i++ {
			fmt.Print(" ")
		}
		fmt.Print(q)
		for i := 1; i <= c; i++ {
			fmt.Print(" ")
		}
	}
}

func CetakJadwal() { //done
	var a, b, c int
	a = len(sch1[0].Judul)
	b = len(sch1[1].Judul)
	c = len(sch1[2].Judul)

	fmt.Println()
	fmt.Println("---------------------------------Jadwal Bioskop BojongSWANG---------------------------------")
	fmt.Println("____________________________________________________________________________________________")
	fmt.Println("|  \\          |                         |                         |                         |")
	fmt.Println("|   \\         |                         |                         |                         |")
	fmt.Println("|    \\ Studio |                         |                         |                         |")
	fmt.Println("|     \\       |                         |                         |                         |")
	fmt.Println("|      \\      |                         |                         |                         |")
	fmt.Println("|       \\     |        Studio  1        |        Studio  2        |        Studio  3        |")
	fmt.Println("|        \\    |                         |                         |                         |")
	fmt.Println("|   Jam   \\   |                         |                         |                         |")
	fmt.Println("|          \\  |                         |                         |                         |")
	fmt.Println("|           \\ |                         |                         |                         |")
	fmt.Println("|____________\\|_________________________|_________________________|_________________________|")
	fmt.Println("|             |                         |                         |                         |")
	fmt.Print("| 10:00-13:00 |")
	SubCetakJadwal(a, sch1[0].Judul)
	fmt.Print("|")
	SubCetakJadwal(b, sch1[1].Judul)
	fmt.Print("|")
	SubCetakJadwal(c, sch1[2].Judul)
	fmt.Println("|")
	fmt.Println("|_____________|_________________________|_________________________|_________________________|")
	fmt.Println("|             |                         |                         |                         |")
	fmt.Print("| 14:00-17:00 |")
	SubCetakJadwal(b, sch2[0].Judul)
	fmt.Print("|")
	SubCetakJadwal(c, sch2[1].Judul)
	fmt.Print("|")
	SubCetakJadwal(a, sch2[2].Judul)
	fmt.Println("|")
	fmt.Println("|_____________|_________________________|_________________________|_________________________|")
	fmt.Println("|             |                         |                         |                         |")
	fmt.Print("| 18:00-21:00 |")
	SubCetakJadwal(c, sch3[0].Judul)
	fmt.Print("|")
	SubCetakJadwal(a, sch3[1].Judul)
	fmt.Print("|")
	SubCetakJadwal(b, sch3[2].Judul)
	fmt.Println("|")
	fmt.Println("|_____________|_________________________|_________________________|_________________________|")
	fmt.Println()

	var pilih int
	fmt.Println("Menu:")
	fmt.Println("[1] Beli tiket")
	fmt.Println("[0] Kembali ke menu awal\n")
	fmt.Print("Pilihan anda: ")
	fmt.Scanln(&pilih)
	if pilih == 1 {
		BeliTiket(&sch1, &sch2, &sch3)
	} else if pilih == 0 {
		MenuAwal()
	}
}

func EditViews(sch1 *SchContent, sch2 *SchContent, sch3 *SchContent) {
	fmt.Print("[1] ", sch1[0].Judul, "\n")
	fmt.Print("[2] ", sch1[1].Judul, "\n")
	fmt.Print("[3] ", sch1[2].Judul, "\n")
	fmt.Print("Masukkan view film yang akan diedit: ")
	var n, exit int
	fmt.Scanln(&n)
	for n != 1 && n != 2 && n != 3 {
		fmt.Println("Tidak ada film dengan nomor ", n, ", mohon input nomor film kembali.")
		fmt.Print("Masukkan view film yang akan diedit: ")
		fmt.Scanln(&n)
	}
	if n == 1 {
		fmt.Print("Ubah view film ", sch1[0].Judul, " pada jadwal 10:00-13:00 menjadi: ")
		fmt.Scanln(&sch1[0].Views)
		fmt.Print("Ubah view film ", sch1[0].Judul, " pada jadwal 14:00-17:00 menjadi: ")
		fmt.Scanln(&sch2[2].Views)
		fmt.Print("Ubah view film ", sch1[0].Judul, " pada jadwal 18:00-21:00 menjadi: ")
		fmt.Scanln(&sch3[1].Views)
	} else if n == 2 {
		fmt.Print("Ubah view film ", sch1[1].Judul, " pada jadwal 10:00-13:00 menjadi: ")
		fmt.Scanln(&sch1[1].Views)
		fmt.Print("Ubah view film ", sch1[1].Judul, " pada jadwal 14:00-17:00 menjadi: ")
		fmt.Scanln(&sch2[0].Views)
		fmt.Print("Ubah view film ", sch1[1].Judul, " pada jadwal 18:00-21:00 menjadi: ")
		fmt.Scanln(&sch3[2].Views)
	} else if n == 3 {
		fmt.Print("Ubah view film ", sch1[2].Judul, " pada jadwal 10:00-13:00 menjadi: ")
		fmt.Scanln(&sch1[2].Views)
		fmt.Print("Ubah view film ", sch1[2].Judul, " pada jadwal 14:00-17:00 menjadi: ")
		fmt.Scanln(&sch2[1].Views)
		fmt.Print("Ubah view film ", sch1[2].Judul, " pada jadwal 18:00-21:00 menjadi: ")
		fmt.Scanln(&sch3[0].Views)
	}

	fmt.Println("Keluar?")
	fmt.Println("[1] Ya")
	fmt.Println("[2] Tidak")
	fmt.Print("Pilihan: ")
	fmt.Scanln(&exit)
	fmt.Println()
	if exit == 1 {
		MenuAdmin()
	} else if exit == 2 {
		EditViews(sch1, sch2, sch3)
	}
}

func EditDurasi(sch1 *SchContent, sch2 *SchContent, sch3 *SchContent) {
	fmt.Print("[1] ", sch1[0].Judul, "\n")
	fmt.Print("[2] ", sch1[1].Judul, "\n")
	fmt.Print("[3] ", sch1[2].Judul, "\n")
	fmt.Print("Masukkan durasi film yang akan diedit: ")
	var n, exit int
	fmt.Scanln(&n)
	for n != 1 && n != 2 && n != 3 {
		fmt.Println("Tidak ada film dengan nomor ", n, ", mohon input nomor film kembali.")
		fmt.Print("Masukkan view film yang akan diedit: ")
		fmt.Scanln(&n)
	}
	if n == 1 {
		fmt.Print("Ubah durasi film \"", sch1[0].Judul, "\" menjadi: ")
		fmt.Scanln(&sch1[0].Durasi)
		sch2[2].Durasi = sch1[0].Durasi
		sch3[1].Durasi = sch1[0].Durasi
	} else if n == 2 {
		fmt.Print("Ubah durasi film \"", sch1[1].Judul, "\" menjadi: ")
		fmt.Scanln(&sch1[1].Durasi)
		sch2[0].Durasi = sch1[0].Durasi
		sch3[2].Durasi = sch1[0].Durasi
	} else if n == 3 {
		fmt.Print("Ubah durasi film \"", sch1[2].Judul, "\" menjadi: ")
		fmt.Scanln(&sch1[2].Durasi)
		sch2[1].Durasi = sch1[0].Durasi
		sch3[0].Durasi = sch1[0].Durasi
	}

	fmt.Println("Keluar?")
	fmt.Println("[1] Ya")
	fmt.Println("[2].Tidak")
	fmt.Print("Pilihan: ")
	fmt.Scanln(&exit)
	fmt.Println()
	if exit == 1 {
		MenuAdmin()
	} else if exit == 2 {
		EditDurasi(sch1, sch2, sch3)
	}
}

func EditFilm() {
	fmt.Println("[1] Edit durasi")
	fmt.Println("[2] Edit views")
	fmt.Println("[0] Kembali ke halaman menu admin")
	fmt.Print("Pilihan: ")
	var pil int
	fmt.Scanln(&pil)
	switch pil {
	case 1:
		EditDurasi(&sch1, &sch2, &sch3)
	case 2:
		EditViews(&sch1, &sch2, &sch3)
	case 0:
		MenuAdmin()
	}
}

func login() { //done
	var pass string
	fmt.Print("\nMasukkan password: ")
	fmt.Scanln(&pass)
	if pass == "bioskopnomorsatu" {
		MenuAdmin()
	} else {
		fmt.Println("Password salah.\nMengalihkan ke menu awal...\n")
		MenuAwal()
	}
}

func DeleteSelectMovie(sch1 *SchContent, sch2 *SchContent, sch3 *SchContent) { //done
	var n int
	fmt.Print("[1] ", sch1[0].Judul, "\n")
	fmt.Print("[2] ", sch1[1].Judul, "\n")
	fmt.Print("[3] ", sch1[2].Judul, "\n")
	fmt.Println()
	fmt.Print("Pilih film yang datanya akan dihapus : ")
	fmt.Scanln(&n)
	for n != 1 && n != 2 && n != 3 {
		fmt.Println("Tidak ada film dengan nomor ", n, ", mohon input nomor film kembali.")
		fmt.Print("Pilih film yang datanya akan dihapus: ")
		fmt.Scanln(&n)
	}
	if n == 1 {
		if Popular[0].Judul == sch1[0].Judul {
			Popular[0].Judul = ""
			Popular[0].Durasi = ""
			Popular[0].Views = 0
		} else if Popular[1].Judul == sch1[0].Judul {
			Popular[1].Judul = ""
			Popular[1].Durasi = ""
			Popular[1].Views = 0
		} else if Popular[2].Judul == sch1[0].Judul {
			Popular[2].Judul = ""
			Popular[2].Durasi = ""
			Popular[2].Views = 0
		}

		sch1[0].Judul = ""
		sch1[0].Durasi = ""
		sch1[0].Views = 0

		sch2[2].Judul = ""
		sch2[2].Durasi = ""
		sch2[2].Views = 0

		sch3[1].Judul = ""
		sch3[1].Durasi = ""
		sch3[1].Views = 0

	} else if n == 2 {
		if Popular[0].Judul == sch2[0].Judul {
			Popular[0].Judul = ""
			Popular[0].Durasi = ""
			Popular[0].Views = 0
		} else if Popular[1].Judul == sch2[0].Judul {
			Popular[1].Judul = ""
			Popular[1].Durasi = ""
			Popular[1].Views = 0
		} else if Popular[2].Judul == sch2[0].Judul {
			Popular[2].Judul = ""
			Popular[2].Durasi = ""
			Popular[2].Views = 0
		}

		sch1[1].Judul = ""
		sch1[1].Durasi = ""
		sch1[1].Views = 0

		sch2[0].Judul = ""
		sch2[0].Durasi = ""
		sch2[0].Views = 0

		sch3[2].Judul = ""
		sch3[2].Durasi = ""
		sch3[2].Views = 0

	} else if n == 3 {
		if Popular[0].Judul == sch3[0].Judul {
			Popular[0].Judul = ""
			Popular[0].Durasi = ""
			Popular[0].Views = 0
		} else if Popular[1].Judul == sch3[0].Judul {
			Popular[1].Judul = ""
			Popular[1].Durasi = ""
			Popular[1].Views = 0
		} else if Popular[2].Judul == sch3[0].Judul {
			Popular[2].Judul = ""
			Popular[2].Durasi = ""
			Popular[2].Views = 0
		}

		sch1[2].Judul = ""
		sch1[2].Durasi = ""
		sch1[2].Views = 0

		sch2[1].Judul = ""
		sch2[1].Durasi = ""
		sch2[1].Views = 0

		sch3[0].Judul = ""
		sch3[0].Durasi = ""
		sch3[0].Views = 0
	}
	MenuAdmin()
}

func DeleteAllMovie(sch1 *SchContent, sch2 *SchContent, sch3 *SchContent) { //done, harga tetap nggak diubah
	var i int
	for i = 0; i < 3; i++ {
		sch1[i].Judul = ""
		sch1[i].Durasi = ""
		sch1[i].Views = 0

		sch2[i].Judul = ""
		sch2[i].Durasi = ""
		sch2[i].Views = 0

		sch3[i].Judul = ""
		sch3[i].Durasi = ""
		sch3[i].Views = 0

		Popular[i].Judul = ""
		Popular[i].Durasi = ""
		Popular[i].Views = 0
	}
	fmt.Println("Seluruh data film sudah terhapus. Mengalihkan ke halaman menu admin...\n")
	MenuAdmin()
}

func InputJudul() string {
	var (
		Title, name string
		TitleLength int
	)
	TitleLength = 0
	fmt.Print("Masukkan judul film: ")
	fmt.Scan(&name)
	for name != "###" && TitleLength <= 25 {
		TitleLength += len(name)
		if TitleLength <= 25 {
			Title = Title + " " + name
			fmt.Scan(&name)
		} else {
			fmt.Println("Judul film melebihi jumlah maksimal karakter. Mohon input ulang.")
			return InputJudul()
		}
	}
	return Title
}

func InputFilm(sch1 *SchContent, sch2 *SchContent, sch3 *SchContent) { //done, cuma ngerubah judul sama durasi, view 0 karena baru didelete, harga tetap untuk semua
	var (
		i int
	)
	for i = 0; i < 3; i++ {
		if sch1[i].Judul == "" {
			fmt.Println("Ada slot film kosong pada slot ke-", i+1)
			/*fmt.Println("Ketentuan input judul: (1) Tambahkan penanda \"###\" pada akhir judul (dipisah dengan spasi)")
			fmt.Println("                       (2) Jumlah karakter maksimal adalah 25 karakter (tidak termasuk penanda)")
			sch1[i].Judul = InputJudul()*/
			fmt.Scanln(&sch1[i].Judul)

			fmt.Print("Masukkan durasi film dalam menit: ")
			fmt.Scanln(&sch1[i].Durasi)
			sch1[i].Durasi += " menit"

			if i == 0 {
				sch2[2] = sch1[0]
				sch3[1] = sch1[0]
			} else if i == 1 {
				sch2[0] = sch1[1]
				sch3[2] = sch1[1]
			} else if i == 2 {
				sch2[1] = sch1[2]
				sch3[0] = sch1[2]
			}
			fmt.Println("Data film pada slot ke-", i+1, " sudah berhasil terisi. Mencari slot film kosong...")
		}
	}
	if i == 3 {
		fmt.Println("Tidak ada slot film kosong. Delete film dan coba lagi.")
	}
	MenuAdmin()
}

func BeliTiket(sch1 *SchContent, sch2 *SchContent, sch3 *SchContent) {
	fmt.Println("Pilihan film:")
	fmt.Print("[1] ", sch1[0].Judul, "\n")
	fmt.Print("[2] ", sch1[1].Judul, "\n")
	fmt.Print("[3] ", sch1[2].Judul, "\n\n")

	var (
		slot, jadwal, idx1, idx2 int
	)
	fmt.Println("Silakan isi form pembelian tiket di bawah ini.")
	fmt.Print("Nomor film: ")
	fmt.Scanln(&slot)
	for slot != 1 && slot != 2 && slot != 3 {
		fmt.Println("Tidak ada film dengan nomor ", slot)
		fmt.Print("Nomor film: ")
		fmt.Scanln(&slot)
	}

	fmt.Println("Pilihan jadwal: [1] 10:00-13:00")
	fmt.Println("                [2] 14:00-17:00")
	fmt.Println("                [3] 18:00-21:00")
	fmt.Print("Nomor jadwal: ")
	fmt.Scanln(&jadwal)
	for jadwal != 1 && jadwal != 2 && jadwal != 3 {
		fmt.Println("Tidak ada film dengan nomor ", jadwal)
		fmt.Print("Nomor jadwal: ")
		fmt.Scanln(&jadwal)
	}

	switch slot {
	case 1:
		if jadwal == 1 {
			idx1, idx2 = FindEmpty(sch1[0])
			if idx1 == 99 {
				fmt.Println("Mohon maaf, tiket film ini pada jadwal yang dipilih sudah habis.")
				MenuAwal()
			} else {
				fmt.Print("Terdapat kursi kosong bernomor ", idx1+1, "-", idx2+1, "\n")
				sch1[0].Kursi[idx1][idx2] = 1
				fmt.Println("Pembelian tiket berhasil. Mengalihkan ke menu awal...")
				sch1[0].Views++
				MenuAwal()
			}
		} else if jadwal == 2 {
			idx1, idx2 = FindEmpty(sch2[2])
			if idx1 == 99 {
				fmt.Println("Mohon maaf, tiket film ini pada jadwal yang dipilih sudah habis.")
				MenuAwal()
			} else {
				fmt.Print("Terdapat kursi kosong bernomor ", idx1+1, "-", idx2+1, "\n")
				sch2[2].Kursi[idx1][idx2] = 1
				fmt.Println("Pembelian tiket berhasil. Mengalihkan ke menu awal...")
				sch2[2].Views++
				MenuAwal()
			}
		} else {
			idx1, idx2 = FindEmpty(sch3[1])
			if idx1 == 99 {
				fmt.Println("Mohon maaf, tiket film ini pada jadwal yang dipilih sudah habis.")
				MenuAwal()
			} else {
				fmt.Print("Terdapat kursi kosong bernomor ", idx1+1, "-", idx2+1, "\n")
				sch3[1].Kursi[idx1][idx2] = 1
				fmt.Println("Pembelian tiket berhasil. Mengalihkan ke menu awal...")
				sch3[1].Views++
				MenuAwal()
			}
		}
	case 2:
		if jadwal == 1 {
			idx1, idx2 = FindEmpty(sch1[1])
			if idx1 == 99 {
				fmt.Println("Mohon maaf, tiket film ini pada jadwal yang dipilih sudah habis.")
				MenuAwal()
			} else {
				fmt.Print("Terdapat kursi kosong bernomor ", idx1+1, "-", idx2+1, "\n")
				sch1[1].Kursi[idx1][idx2] = 1
				fmt.Println("Pembelian tiket berhasil. Mengalihkan ke menu awal...")
				sch1[1].Views++
				MenuAwal()
			}
		} else if jadwal == 2 {
			idx1, idx2 = FindEmpty(sch2[0])
			if idx1 == 99 {
				fmt.Println("Mohon maaf, tiket film ini pada jadwal yang dipilih sudah habis.")
				MenuAwal()
			} else {
				fmt.Print("Terdapat kursi kosong bernomor ", idx1+1, "-", idx2+1, "\n")
				sch2[0].Kursi[idx1][idx2] = 1
				fmt.Println("Pembelian tiket berhasil. Mengalihkan ke menu awal...")
				sch2[0].Views++
				MenuAwal()
			}
		} else {
			idx1, idx2 = FindEmpty(sch3[2])
			if idx1 == 99 {
				fmt.Println("Mohon maaf, tiket film ini pada jadwal yang dipilih sudah habis.")
				MenuAwal()
			} else {
				fmt.Print("Terdapat kursi kosong bernomor ", idx1+1, "-", idx2+1, "\n")
				sch3[2].Kursi[idx1][idx2] = 1
				fmt.Println("Pembelian tiket berhasil. Mengalihkan ke menu awal...")
				sch3[2].Views++
				MenuAwal()
			}
		}
	case 3:
		if jadwal == 1 {
			idx1, idx2 = FindEmpty(sch1[2])
			if idx1 == 99 {
				fmt.Println("Mohon maaf, tiket film ini pada jadwal yang dipilih sudah habis.")
				MenuAwal()
			} else {
				fmt.Print("Terdapat kursi kosong bernomor ", idx1+1, "-", idx2+1, "\n")
				sch1[2].Kursi[idx1][idx2] = 1
				fmt.Println("Pembelian tiket berhasil. Mengalihkan ke menu awal...")
				sch1[2].Views++
				MenuAwal()
			}
		} else if jadwal == 2 {
			idx1, idx2 = FindEmpty(sch2[1])
			if idx1 == 99 {
				fmt.Println("Mohon maaf, tiket film ini pada jadwal yang dipilih sudah habis.")
				MenuAwal()
			} else {
				fmt.Print("Terdapat kursi kosong bernomor ", idx1+1, "-", idx2+1, "\n")
				sch2[1].Kursi[idx1][idx2] = 1
				fmt.Println("Pembelian tiket berhasil. Mengalihkan ke menu awal...")
				sch2[1].Views++
				MenuAwal()
			}
		} else {
			idx1, idx2 = FindEmpty(sch3[0])
			if idx1 == 99 {
				fmt.Println("Mohon maaf, tiket film ini pada jadwal yang dipilih sudah habis.")
				MenuAwal()
			} else {
				fmt.Print("Terdapat kursi kosong bernomor ", idx1+1, "-", idx2+1, "\n")
				sch3[0].Kursi[idx1][idx2] = 1
				fmt.Println("Pembelian tiket berhasil. Mengalihkan ke menu awal...")
				sch3[0].Views++
				MenuAwal()
			}
		}
	}
}

func MenuAwal() {
	fmt.Println("\n--------------------------------------------------------------------------------------------")
	fmt.Printf("\nMENU : \n")
	fmt.Println("[1] Lihat jadwal")
	fmt.Println("[2] Beli tiket")
	fmt.Println("[3] Login admin")
	fmt.Println("[0] Keluar\n")

	var choice int
	fmt.Print("Masukkan pilihan Anda : ")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		CetakJadwal()
	case 2:
		BeliTiket(&sch1, &sch2, &sch3)
	case 3:
		login()
	case 0:
		fmt.Println()
	}
}

func MenuAdmin() {
	var (
		pil int
	)
	fmt.Println()
	fmt.Println("-----------Menu Admin-----------")
	fmt.Println("[1] Hapus Data Film ")
	fmt.Println("[2] Hapus Semua Data Film")
	fmt.Println("[3] Edit Data Film")
	fmt.Println("[4] Input Data Film")
	fmt.Println("[5] Statistik Penjualan Tiket")
	fmt.Println("[6] Reset penjualan hari ini")
	fmt.Println("[0] Kembali Ke Menu Awal")
	fmt.Println()
	fmt.Print("Masukkan Pilihan Anda : ")
	fmt.Scanln(&pil)
	if pil == 1 {
		DeleteSelectMovie(&sch1, &sch2, &sch3)
	} else if pil == 2 {
		DeleteAllMovie(&sch1, &sch2, &sch3)
	} else if pil == 3 {
		EditFilm()
	} else if pil == 4 {
		InputFilm(&sch1, &sch2, &sch3)
	} else if pil == 5 {
		Report()
	} else if pil == 6 {
		ResetAllSeat(&sch1, &sch2, &sch3)
	} else if pil == 0 {
		MenuAwal()
	}
}

func IsiDefault(sch1 *SchContent, sch2 *SchContent, sch3 *SchContent) {
	sch1[0].Judul = "Joker"
	sch1[0].Durasi = "122 menit"
	sch1[0].Harga = 20000

	sch1[1].Judul = "Toy Story 4"
	sch1[1].Durasi = "100 menit"
	sch1[1].Harga = 20000

	sch1[2].Judul = "Avengers: Endgame"
	sch1[2].Durasi = "182 menit"
	sch1[2].Harga = 20000
	//---------------------------------
	sch2[0].Judul = "Toy Story 4"
	sch2[0].Durasi = "100 menit"
	sch2[0].Harga = 20000

	sch2[1].Judul = "Avengers: Endgame"
	sch2[1].Durasi = "182 menit"
	sch2[1].Harga = 20000

	sch2[2].Judul = "Joker"
	sch2[2].Durasi = "122 menit"
	sch2[2].Harga = 20000
	//---------------------------------
	sch3[0].Judul = "Avengers: Endgame"
	sch3[0].Durasi = "182 menit"
	sch3[0].Harga = 20000

	sch3[1].Judul = "Joker"
	sch3[1].Durasi = "122 menit"
	sch3[1].Harga = 20000

	sch3[2].Judul = "Toy Story 4"
	sch3[2].Durasi = "100 menit"
	sch3[2].Harga = 20000

	for i := 0; i < 3; i++ {
		ResetSeat(&sch1[i].Kursi)
		ResetSeat(&sch2[i].Kursi)
		ResetSeat(&sch3[i].Kursi)
	}
}

func main() {
	IsiDefault(&sch1, &sch2, &sch3)

	fmt.Println("\n                                    BIOSKOP  BOJONGSWANG")
	fmt.Println("                              Jalan Terusan Buah Batu, Bandung")

	MenuAwal()

	fmt.Println("---------------------Terima kasih telah mengunjungi Bioskop BojongSWANG---------------------")
	fmt.Println("                             Kami tunggu kedatangannya kembali")

}
