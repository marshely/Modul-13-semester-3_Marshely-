package main

import "fmt"

type Buku struct {
	id 			int
	judul 		string
	penulis 	string
	penerbit 	string
	eksemplar 	int
	tahun 		int
	rating 		float64
}

type DaftarBuku []Buku

func DaftarDataBuku(pustaka *DaftarBuku, n *int) {
	fmt.Println("Masukkan data buku:")
	for i := 0; i < *n; i++ {
		var buku Buku
		fmt.Printf("Buku %d\n", i+1)
		fmt.Print("ID: ")
		fmt.Scan(&buku.id)
		fmt.Print("Judul: ")
		fmt.Scan(&buku.judul)
		fmt.Print("Penulis: ")
		fmt.Scan(&buku.penulis)
		fmt.Print("Penerbit: ")
		fmt.Scan(&buku.penerbit)
		fmt.Print("Eksemplar: ")
		fmt.Scan(&buku.eksemplar)
		fmt.Print("Tahun: ")
		fmt.Scan(&buku.tahun)
		fmt.Print("Rating: ")
		fmt.Scan(&buku.rating)
		*pustaka = append(*pustaka, buku)

		fmt.Println()
	}
}
func CetakFavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		fmt.Println("Tidak ada buku dalam pustaka.")
		return
	}
	terfavorit := pustaka[0]
	for i := 1; i < n; i++ {
		if pustaka[i].rating > terfavorit.rating {
			terfavorit = pustaka[i]
		}
	}
	fmt.Println("\nBuku Terfavorit:")
	fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %.2f\n", terfavorit.judul, terfavorit.penulis, terfavorit.penerbit, terfavorit.tahun, terfavorit.rating)
}

func UrutanBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := (*pustaka)[i]
		j := i - 1
		for j >= 0 && (*pustaka)[j].rating < key.rating {
			(*pustaka)[j+1] = (*pustaka)[j]
			j--
		}
		(*pustaka)[j+1] = key
	}
}

func CetakTerbaru(pustaka DaftarBuku, n int) {
	fmt.Println("\n5 Buku dengan Rating Tertinggi:")
	limit := 5
	if n < 5 {
		limit = n
	}
	for i := 0; i < limit; i++ {
		fmt.Printf("%d. %s (Rating: %.2f)\n", i+1, pustaka[i].judul,pustaka[i].rating)
 	}
}
func CariBuku(pustaka DaftarBuku, n int, r float64) {
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if pustaka[mid].rating == r {
			fmt.Println("\nBuku ditemukan:")
			fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun:%d, Eksemplar: %d, Rating: %.2f\n",
		pustaka[mid].judul, pustaka[mid].penulis, pustaka[mid].penerbit, pustaka[mid].tahun, pustaka[mid].eksemplar, pustaka[mid].rating)
		return
		} else if pustaka[mid].rating < r {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	fmt.Println("\nTidak ada buku dengan rating seperti itu.")
}

func main() {
	var pustaka DaftarBuku
	var n int

	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(&n)

		DaftarDataBuku(&pustaka, &n)

		CetakFavorit(pustaka, n)

		UrutanBuku(&pustaka, n)

		CetakTerbaru(pustaka, n)

		var r float64
		fmt.Print("\nMasukkan rating buku yang ingin dicari: ")
		fmt.Scan(&r)
		CariBuku(pustaka, n, r)
	}
		