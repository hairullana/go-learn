package main

import (
	"fmt"
	"os"
)

type Alamat struct {
	kelurahan string
	kecamatan string
	kota      string
	provinsi  string
}

type Person struct {
	nama      string
	alamat    Alamat
	pekerjaan string
	alasan    string
}

func main() {
	var Person = []Person{
		{
			nama: "Hairul Lana",
			alamat: Alamat{
				kelurahan: "Sengkidu",
				kecamatan: "Manggis",
				kota:      "Karangasem",
				provinsi:  "Bali",
			},
			pekerjaan: "Software Engineer",
			alasan:    "Karena ingin mempelajari bahasa baru",
		},
		{
			nama: "Hairul Lana",
			alamat: Alamat{
				kelurahan: "Sengkidu",
				kecamatan: "Manggis",
				kota:      "Karangasem",
				provinsi:  "Bali",
			},
			pekerjaan: "Software Engineer",
			alasan:    "Karena ingin mempelajari bahasa baru",
		},
		{
			nama: "Hairul Lana",
			alamat: Alamat{
				kelurahan: "Sengkidu",
				kecamatan: "Manggis",
				kota:      "Karangasem",
				provinsi:  "Bali",
			},
			pekerjaan: "Software Engineer",
			alasan:    "Karena ingin mempelajari bahasa baru",
		},
		{
			nama: "Hairul Lana",
			alamat: Alamat{
				kelurahan: "Sengkidu",
				kecamatan: "Manggis",
				kota:      "Karangasem",
				provinsi:  "Bali",
			},
			pekerjaan: "Software Engineer",
			alasan:    "Karena ingin mempelajari bahasa baru",
		},
		{
			nama: "Hairul Lana",
			alamat: Alamat{
				kelurahan: "Sengkidu",
				kecamatan: "Manggis",
				kota:      "Karangasem",
				provinsi:  "Bali",
			},
			pekerjaan: "Software Engineer",
			alasan:    "Karena ingin mempelajari bahasa baru",
		},
		{
			nama: "Hairul Lana",
			alamat: Alamat{
				kelurahan: "Sengkidu",
				kecamatan: "Manggis",
				kota:      "Karangasem",
				provinsi:  "Bali",
			},
			pekerjaan: "Software Engineer",
			alasan:    "Karena ingin mempelajari bahasa baru",
		},
		{
			nama: "Hairul Lana",
			alamat: Alamat{
				kelurahan: "Sengkidu",
				kecamatan: "Manggis",
				kota:      "Karangasem",
				provinsi:  "Bali",
			},
			pekerjaan: "Software Engineer",
			alasan:    "Karena ingin mempelajari bahasa baru",
		},
		{
			nama: "Hairul Lana",
			alamat: Alamat{
				kelurahan: "Sengkidu",
				kecamatan: "Manggis",
				kota:      "Karangasem",
				provinsi:  "Bali",
			},
			pekerjaan: "Software Engineer",
			alasan:    "Karena ingin mempelajari bahasa baru",
		},
		{
			nama: "Hairul Lana",
			alamat: Alamat{
				kelurahan: "Sengkidu",
				kecamatan: "Manggis",
				kota:      "Karangasem",
				provinsi:  "Bali",
			},
			pekerjaan: "Software Engineer",
			alasan:    "Karena ingin mempelajari bahasa baru",
		},
		{
			nama: "Hairul Lana",
			alamat: Alamat{
				kelurahan: "Sengkidu",
				kecamatan: "Manggis",
				kota:      "Karangasem",
				provinsi:  "Bali",
			},
			pekerjaan: "Software Engineer",
			alasan:    "Karena ingin mempelajari bahasa baru",
		},
	}

	if os.Args[0] {
		var index int = os.Args[0]
		showPersonData(index)
	}
}

func showPersonData(index int) {
	fmt.Println("Nama: ", Person)
}
