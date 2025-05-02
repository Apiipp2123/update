package view

import (
	"bufio"
	"fmt"
	controller "node/controllers"
	"node/model"
	"os"
	"strconv"
	"strings"
)

func InputKomik() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Judul: ")
	judul, _ := reader.ReadString('\n')

	fmt.Print("Tahun Terbit: ")
	tahunStr, _ := reader.ReadString('\n')
	tahunTerbit, _ := strconv.Atoi(strings.TrimSpace(tahunStr))

	fmt.Print("Status: ")
	status, _ := reader.ReadString('\n')

	fmt.Print("Rating: ")
	ratingStr, _ := reader.ReadString('\n')
	rating, _ := strconv.ParseFloat(strings.TrimSpace(ratingStr), 64)

	fmt.Print("Genre: ")
	genre, _ := reader.ReadString('\n')

	fmt.Println("=== Data Penulis ===")
	fmt.Print("Nama: ")
	pNama, _ := reader.ReadString('\n')
	fmt.Print("Email: ")
	pEmail, _ := reader.ReadString('\n')
	fmt.Print("Asal: ")
	pAsal, _ := reader.ReadString('\n')
	fmt.Print("Tahun Lahir: ")
	pLahirStr, _ := reader.ReadString('\n')
	pLahir, _ := strconv.Atoi(strings.TrimSpace(pLahirStr))

	fmt.Println("=== Data Ilustrator ===")
	fmt.Print("Nama: ")
	iNama, _ := reader.ReadString('\n')
	fmt.Print("Email: ")
	iEmail, _ := reader.ReadString('\n')
	fmt.Print("Asal: ")
	iAsal, _ := reader.ReadString('\n')
	fmt.Print("Tahun Lahir: ")
	iLahirStr, _ := reader.ReadString('\n')
	iLahir, _ := strconv.Atoi(strings.TrimSpace(iLahirStr))

	fmt.Println("=== Data Penerbit ===")
	fmt.Print("Nama: ")
	pbNama, _ := reader.ReadString('\n')
	fmt.Print("Alamat: ")
	pbAlamat, _ := reader.ReadString('\n')
	fmt.Print("Tahun Berdiri: ")
	pbTahunStr, _ := reader.ReadString('\n')
	pbTahun, _ := strconv.Atoi(strings.TrimSpace(pbTahunStr))
	fmt.Print("Website: ")
	pbWeb, _ := reader.ReadString('\n')

	komik := model.NewKomik(
		controller.GetNextID(),
		tahunTerbit,
		strings.TrimSpace(judul),
		strings.TrimSpace(status),
		strings.TrimSpace(genre),
		strings.TrimSpace(pNama),
		strings.TrimSpace(pEmail),
		strings.TrimSpace(pAsal),
		pLahir,
		strings.TrimSpace(iNama),
		strings.TrimSpace(iEmail),
		strings.TrimSpace(iAsal),
		iLahir,
		strings.TrimSpace(pbNama),
		strings.TrimSpace(pbAlamat),
		strings.TrimSpace(pbWeb),
		pbTahun,
		rating,
	)

	controller.TambahKomik(komik)
	fmt.Println("Komik berhasil ditambahkan!")
}
