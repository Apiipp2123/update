package model

import (
	"node/database"
	"strings"
)

func NewKomik(
	id, tahunTerbit int,
	judul, status, genre, penulisNama, penulisEmail, penulisAsal string, penulisLahir int,
	ilustratorNama, ilustratorEmail, ilustratorAsal string, ilustratorLahir int,
	penerbitNama, alamatPenerbit, websitePenerbit string, tahunBerdiri int,
	rating float64,
) database.Komik {
	return database.Komik{
		ID:          id,
		Judul:       strings.TrimSpace(judul),
		TahunTerbit: tahunTerbit,
		Status:      strings.TrimSpace(status),
		Rating:      rating,
		Genre: database.Genre{
			ID:   id * 10,
			Nama: strings.TrimSpace(genre),
		},
		Penulis: database.Penulis{
			ID:         id * 11,
			Nama:       strings.TrimSpace(penulisNama),
			Email:      strings.TrimSpace(penulisEmail),
			Asal:       strings.TrimSpace(penulisAsal),
			TahunLahir: penulisLahir,
		},
		Ilustrator: database.Ilustrator{
			ID:         id * 12,
			Nama:       strings.TrimSpace(ilustratorNama),
			Email:      strings.TrimSpace(ilustratorEmail),
			Asal:       strings.TrimSpace(ilustratorAsal),
			TahunLahir: ilustratorLahir,
		},
		Penerbit: database.Penerbit{
			ID:           id * 13,
			Nama:         strings.TrimSpace(penerbitNama),
			Alamat:       strings.TrimSpace(alamatPenerbit),
			TahunBerdiri: tahunBerdiri,
			Website:      strings.TrimSpace(websitePenerbit),
		},
	}
}
