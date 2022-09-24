package main
// Hal wajib yang harus dilakukan ketika menggunakan bahasa Go 

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)
/*
	Pada kodingan di atas kita mengimport beberapa package dalam Go
	untuk keperluan dalam project sevice CLI ini. Salah satunya adalah
	package os yang digunakan untuk memberikan argumen yang disisipkan
	ketika mengeksekusi program
*/

type Biodata struct {
	absen     string
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

/*
	Di atas merupakan struct dengan nama Biodata yang dimana terdapat
	field-field untuk keperluan project ini. Adapun field tersebut yaitu
	absen, nama, alamat, pekerjaan, dan alasan. 
*/

func getBiodata(s []Biodata, no_urut int) {
	
	if no_urut <= 5 {
		fmt.Println()
		fmt.Println(strings.Repeat("=", 50))
		
		fmt.Printf("Nomor\t\t: %s\n", s[no_urut-1].absen)
		fmt.Printf("Nama\t\t: %s\n", s[no_urut-1].nama)
		fmt.Printf("Alamat\t\t: %s\n", s[no_urut-1].alamat)
		fmt.Printf("Pekerjaan\t: %s\n", s[no_urut-1].pekerjaan)
		fmt.Printf("Alasan\t\t: %s\n", s[no_urut-1].alasan)

		fmt.Println(strings.Repeat("=", 50))
		fmt.Println()
	} else {
		fmt.Println()
		fmt.Println("Data tidak ditemukan")
		fmt.Println()
	}
}

/*
	Pertama, membuat func getBiodata() dengan parameternya yaitu sebuah 
	slice of struct dari Biodata yang kita berikan pada variabel s. Lalu
	parameter kedua adalah no_urut dengan tipe data integer.

	func getBiodata() ini akan mengambil nilai dari field-field dari slice
	of struct pada fungsi main() di bawah yang disimpan ke dalam variabel 
	students.

	Pada func getBiodata() kita langsung berikan sebuah kondisi, jika nilai
	dari variabel no_urut itu <= 5 maka akan memunculkan output field-field
	pada variabel students dengan cara memanggil parameter pertama yaitu s
	diikuti nilai dari variabel no_urut yang dikurangi dengan 1 lalu pemanggilan
	field pada slice students.

	Jika nilai dari variabel no_urut > 5, maka akan menampilkan pesan "Data tidak
	ditemukan" karena memang kita hanya menggunakan 5 data saja.
*/


func main() {
	no_urut, err := strconv.Atoi(os.Args[1])
	_ = err

	/*
		Kodingan di atas digunakan untuk konversi data dari tipe string ke int 
		dengan nilainya yaitu dari os.Args, hasil konversi ini akan disimpan
		pada variabel no_urut dan err.
	*/
	var students = []Biodata{
		{absen: "1", nama: "Abdul", alamat: "Indonesia", pekerjaan: "front end developer", alasan: "Ingin mengetahui lebih mendalam"},
		{absen: "2", nama: "Hanif", alamat: "Indonesia", pekerjaan: "back end developer", alasan: "Memperbagus portofolio"},
		{absen: "3", nama: "Pratama", alamat: "Indonesia", pekerjaan: "android developer", alasan: "Cari kesibukan aja"},
		{absen: "4", nama: "Ujang", alamat: "Indonesia", pekerjaan: "ios developer", alasan: "Menjalin relasi dengan sesama"},
		{absen: "5", nama: "Udin", alamat: "Indonesia", pekerjaan: "devops engineer", alasan: "Menambah wawasan seputar Go"},
	}

	/*
		Kemudian di atas merupakan kodingan pembuatas slice of struct yang disimpan 
		dalam variabel student. Dimana terdapat 5 data dengan fieldnya yaitu absen,
		nama, alamat, pekerjaan, dan alasan.
	*/
	
	getBiodata(students, no_urut)
	/*
		Dan ini merupakan pemanggilan dari func getBiodata yang diisi dengan parameter
		students dan no_urut. Pemanggilan func ini tentunya untuk melihat dari hasil 
		output pada program service CLI ini.
	*/
}