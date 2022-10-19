package main

import (
	"fmt"
	"os"
	"strconv"
)

// deklarasi struct Alamat
type Alamat struct {
	kelurahan string
	kecamatan string
	kota      string
	provinsi  string
}

// deklarasi struct Person
type Person struct {
	nama      string
	alamat    Alamat
	pekerjaan string
	alasan    string
}

// set data ke persons dari struct Person
var persons = []Person{
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
		nama: "Surya Anto",
		alamat: Alamat{
			kelurahan: "Manggarai",
			kecamatan: "Manggarai",
			kota:      "Jakarta Selatan",
			provinsi:  "Jakarta",
		},
		pekerjaan: "Barista",
		alasan:    "Karena ingin mempelajari pemrograman",
	},
	{
		nama: "Firdaus Zulzul",
		alamat: Alamat{
			kelurahan: "Kebun Raya",
			kecamatan: "Kebun Raya",
			kota:      "Surabaya",
			provinsi:  "Jawa Timur",
		},
		pekerjaan: "IT Support",
		alasan:    "Ingin switch karir jadi IT Programmer",
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
		nama: "Karel Leo Arnaldo",
		alamat: Alamat{
			kelurahan: "Kaliwangi",
			kecamatan: "Kaliwangi",
			kota:      "Banyuwangi",
			provinsi:  "Jawa Timur",
		},
		pekerjaan: "Belum ada",
		alasan:    "Ingin berhenti merokok",
	},
	{
		nama: "Indah Sari",
		alamat: Alamat{
			kelurahan: "Buntu Utara",
			kecamatan: "Buntu Utara",
			kota:      "Cirebon",
			provinsi:  "Banten",
		},
		pekerjaan: "Marketing",
		alasan:    "Bosan menjadi marketing",
	},
	{
		nama: "Karlina Surya Atmaja",
		alamat: Alamat{
			kelurahan: "Wongso",
			kecamatan: "Wongso Utama",
			kota:      "Yogyakarta",
			provinsi:  "DIY",
		},
		pekerjaan: "Wirausaha",
		alasan:    "Jadi programmer karena jualan sepi",
	},
	{
		nama: "Wina Artha Setiawati",
		alamat: Alamat{
			kelurahan: "Mertani",
			kecamatan: "Mertani",
			kota:      "Aceh Barat",
			provinsi:  "Aceh",
		},
		pekerjaan: "Pengusaha",
		alasan:    "Mengisi waktu luang",
	},
	{
		nama: "Muhammad Firdaus Suhartana",
		alamat: Alamat{
			kelurahan: "Mangojo",
			kecamatan: "Manggojo",
			kota:      "Flores",
			provinsi:  "NTT",
		},
		pekerjaan: "Petani Cabai",
		alasan:    "Ingin pekerjaan baru, lahan saya kering",
	},
	{
		nama: "Ahmad Rivaldo Atmaja",
		alamat: Alamat{
			kelurahan: "Dukumu",
			kecamatan: "Dukumu",
			kota:      "Jayapura",
			provinsi:  "Papua",
		},
		pekerjaan: "Atlet Sepak Bola",
		alasan:    "Mengisi waktu liburan saya",
	},
}

// main function
func main() {
	if len(os.Args) > 1 { // jika user mengisi argumen
		// ambil argumen indeks 1
		var index, _ = strconv.Atoi(os.Args[1])
		if index >= len(persons) { // jika argumen indeks lebih dari panjang persons, return pesan ...
			fmt.Println("Data siswa hanya berisi 10 data saja, isi argumen 1 -", len(persons)-1)
		} else { // tampilkan data siswa sesuai indeks
			showPersonData(index)
		}
	} else { // jika user tidak mengisi argumen
		fmt.Println("Berikan argumen / parameter indeks data siswa")
	}
}

// fungsi menampilkan data siswa sesuai indeks
func showPersonData(index int) {
	fmt.Println("Nama: ", persons[index].nama)
	fmt.Printf("Alamat: %s, %s, %s, %s\n", persons[index].alamat.kelurahan, persons[index].alamat.kecamatan, persons[index].alamat.kota, persons[index].alamat.provinsi)
	fmt.Println("Pekerjaan: ", persons[index].pekerjaan)
	fmt.Println("Alasan Belajar Golang: ", persons[index].alasan)
}
