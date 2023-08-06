package main

import "fmt"

// func main() {
// 	// inisialisasi variabel dengan menggunakan tipe data Hitung interface
// 	var bangunDatar intf.Hitung = strcmtd.Lingkaran{Diameter: 20.0}

// 	// menggunakan method yang sudah dibuat
// 	// implementasi lingkaran
// 	bangunDatar = strcmtd.Lingkaran{Diameter: 14.0}
// 	fmt.Println("=====> lingkaran")

// 	// karena method Jarijari() tidak tersedia di interface intf.Hitung, maka untuk mengakses-nya harus di-casting terlebih dahulu ke tipe variabel asli nya (pada kasus ini bertpe strcmtd.Lingkaran)
// 	// setelah itu method nya baru bisa di akses
// 	// untuk melakukan casting interface yaitu dengan menuliskan nama tipe tujuan dalam tanda kurung
// 	// ditempatkan setelah nama interface dengan menggunakan notasi titik (seperti cara megakses field pada sebuah struct, hanya saja ada tanda kurungnya)
// 	fmt.Printf("\n\tjari-jari\t: %.3f\n", bangunDatar.(strcmtd.Lingkaran).Jarijari())
// 	fmt.Printf("\tluas\t\t: %.3f\n", bangunDatar.Luas())
// 	fmt.Printf("\tkeliling\t: %.3f\n\n", bangunDatar.Keliling())

// 	// implementasi persegi
// 	bangunDatar = strcmtd.Persegi{Sisi: 10.0}
// 	fmt.Println("=====> persegi")
// 	fmt.Printf("\n\tluas\t\t: %.3f\n", bangunDatar.Luas())
// 	fmt.Printf("\tkeliling\t: %.3f\n", bangunDatar.Keliling())
// }

// implementasi embeded interface yang telah dibuat
// func main() {
// 	// variabel yang dibuat menggunakan struct yang memiliki method pointer,
// 	// jika di tampung dalam variabel interface harus diambil referensinya terlebih dahulu
// 	var bangunRuang intf.Hitung = &strcmtd.Kubus{Sisi: 4}

// 	fmt.Println("=====> Kubus")
// 	fmt.Println("luas 		: ", bangunRuang.Luas())
// 	fmt.Println("keliling 	: ", bangunRuang.Keliling())
// 	fmt.Println("volume 		: ", bangunRuang.Volume())
// }

// menggunakan interface kosong (any)
// func main() {
// 	// jika sebuah interface ditambahkan kurung kurawal di belakangnya
// 	// maka kegunaannya berpindah dari keyword untuk pembuatan interface menjadi tipe data
// 	var secret interface{} // default value nya adalah nil

// 	fmt.Println("tipe data : ", reflect.TypeOf(secret), " value : ", secret)
// 	secret = "fani alfi"

// 	fmt.Println()
// 	fmt.Println("tipe data : ", reflect.TypeOf(secret), " value : ", secret)

// 	fmt.Println()
// 	secret = []string{"fani", "alfi"}
// 	fmt.Println("tipe data : ", reflect.TypeOf(secret), " value : ", secret)

// 	fmt.Println()
// 	secret = 3.14
// 	fmt.Println("tipe data : ", reflect.TypeOf(secret), " value : ", secret)

// 	fmt.Println()
// 	var data = map[string]interface{}{
// 		"nama": "fani alfirdaus",
// 		"alamat": map[string]string{
// 			"negara":    "indonesia",
// 			"provinsi":  "jawa tengah",
// 			"kabupaten": "karanganyar",
// 			"kecamatan": "matesih",
// 			"desa":      "koripan",
// 			"dusun":     "dukuh",
// 		},
// 		"hobi": []string{"membaca", "ngoding", "belajar"},
// 		"umur": 17,
// 	}

// 	for key, value := range data {
// 		fmt.Printf("key-%s\t: %v\n", key, value)
// 	}

// 	// alias menggunakan interface{}
// 	fmt.Printf("\nalias menggunakan interface\n")
// 	data = map[string]any{
// 		"name": "fanialfi",
// 		"umur": 17,
// 		"alamat": map[string]string{
// 			"kab": "karanganyar",
// 			"kec": "matesih",
// 			"kel": "koripan",
// 			"ds":  "dukuh",
// 		},
// 		"fruits": []string{"apel", "mangga", "pisang", "anggur"},
// 	}

// 	for key, value := range data {
// 		fmt.Printf("key-%s\t: %v\n", key, value)

// 		if key == "alamat" {
// 			// fmt.Println("tipe value alamat ", reflect.TypeOf(value)) // digunakan untuk mengecek tipe data value

// 			// value.(map[string]string) digunakan untuk melakukan casting tipe interface{} ( interface kosong) ke tipe data asli-nya
// 			for key, value := range value.(map[string]string) {
// 				fmt.Printf("%s\t: %s\n", key, value)
// 			}
// 		} else if key == "fruits" {
// 			// value dari value data di casting lagi dari interface{} menjadi bentuk aslinya ([]string)
// 			// strings.Join() digunakan untuk menggabungkan slice yang tipe data,nya string
// 			// dimana parameter pertama adalah sebuah element bertipe slice string yang akan dijoin
// 			// dan parameter kedua adalah penghubungnya
// 			fmt.Println(strings.Join(value.([]string), ", "), "adalah buah kesauaan saya")
// 		}
// 	}
// }

// {
// 	// casting interface kosong ke variabel pointer
// func main() {
// 	type person []string

// 	// variabel data dideklarasikan menggunakan interface kosong
// 	// diisi dengan slice string pointer
// 	var data interface{} = &person{"fani", "alfi"}

// 	// cara casting interface any ke slice string pointer sama ketika casting dari interface any ke bentuk aslinya
// 	// yang membedakan adalah dengan menambahkan tanda asterisk (*) didepan tipe data aslinya
// 	// data.(*person)
// 	nama := data.(*person)
// 	fmt.Println("nama saya adalah", strings.Join(*nama, ""))
// 	}
// }

// kombinasi slice map dan interface{}
func main() {
	var persons = []map[string]any{
		{"name": "fani alfi", "umur": 17},
		{"name": "wick", "umur": 18},
		{"name": "stevan", "umur": 19},
	}

	for _, person := range persons {
		fmt.Println(person["name"], "age is", person["umur"])
	}

	fmt.Printf("\nmemanfaatkan slice dan interface{} untuk membuat data array yang isinya bisa apa saja\n\n")

	var fruits = []any{
		map[string]any{"name": "strawberry", "total": 10},
		[]string{"manggo", "pineapple", "papaya"},
		"orange",
	}

	for _, fruit := range fruits {
		fmt.Println(fruit)
	}
}
