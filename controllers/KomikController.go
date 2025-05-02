package controller

import (
	"fmt"
	"node/database"
	"strings"
)

var KomikList []database.Komik
var komikIDCounter = 1

func GetNextID() int {
	return komikIDCounter
}

func TambahKomik(komik database.Komik) {
	komik.ID = komikIDCounter
	KomikList = append(KomikList, komik)
	komikIDCounter++
}


func TampilkanSemuaKomik() {
	if len(KomikList) == 0 {
		fmt.Println("Belum ada data komik.")
		return
	}
	for _, k := range KomikList {
		fmt.Println("ID: ", k.ID)
		fmt.Println("Judul: ", k.Judul)
		fmt.Println("Penulis: ", k.Penulis.Nama, k.Penulis.Email)
		fmt.Println("Ilustrator: ", k.Ilustrator.Nama, k.Ilustrator.Email)
		fmt.Println("Tahun Terbit: ",  k.TahunTerbit)
		fmt.Println("Penerbit: ",  k.Penerbit.Nama, k.Penerbit.Alamat)
		fmt.Println("Status: ", k.Status)
		fmt.Println("Genre: ", k.Genre.Nama)
		fmt.Println("Rating: ", k.Rating)
	}
}

func UpdateKomik(id int, judulBaru string) {
	for i := range KomikList {
		if KomikList[i].ID == id {
			KomikList[i].Judul = strings.TrimSpace(judulBaru)
			fmt.Println("Komik berhasil diperbarui.")
			return
		}
	}
	fmt.Println("Komik tidak ditemukan.")
}

func DeleteKomik(id int) {
	for i := range KomikList {
		if KomikList[i].ID == id {
			KomikList = append(KomikList[:i], KomikList[i+1:]...)
			fmt.Println("Komik berhasil dihapus.")
			return
		}
	}
	fmt.Println("Komik tidak ditemukan.")
}

func CariKomikByJudul(keyword string) {
	keyword = strings.ToLower(keyword)
	found := false
	for _, k := range KomikList {
		if strings.Contains(strings.ToLower(k.Judul), keyword) {
			fmt.Println("ID: ", k.ID)
			fmt.Println("Judul: ", k.Judul)
			fmt.Println("Penulis: ", k.Penulis.Nama, k.Penulis.Email)
			fmt.Println("Ilustrator: ", k.Ilustrator.Nama, k.Ilustrator.Email)
			fmt.Println("Tahun Terbit: ",  k.TahunTerbit)
			fmt.Println("Penerbit: ",  k.Penerbit.Nama, k.Penerbit.Alamat)
			fmt.Println("Status: ", k.Status)
			fmt.Println("Genre: ", k.Genre.Nama)
			fmt.Println("Rating: ", k.Rating)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ditemukan.")
	}
}

func CariKomikByGenre(keyword string) {
	keyword = strings.ToLower(keyword)
	found := false
	for _, k := range KomikList {
		if strings.Contains(strings.ToLower(k.Genre.Nama), keyword) {
			fmt.Println("ID: ", k.ID)
			fmt.Println("Judul: ", k.Judul)
			fmt.Println("Penulis: ", k.Penulis.Nama, k.Penulis.Email)
			fmt.Println("Ilustrator: ", k.Ilustrator.Nama, k.Ilustrator.Email)
			fmt.Println("Tahun Terbit: ",  k.TahunTerbit)
			fmt.Println("Penerbit: ",  k.Penerbit.Nama, k.Penerbit.Alamat)
			fmt.Println("Status: ", k.Status)
			fmt.Println("Genre: ", k.Genre.Nama)
			fmt.Println("Rating: ", k.Rating)
		}
	}
	if !found {
		fmt.Println("Tidak ditemukan.")
	}
}
