package main

import (
	"fmt"
	"os"
	"strconv"
)

type Biodata struct {
	name    string
	address string
	job     string
	reason  string
}

func main() {
	argsRaw := os.Args

	// DATA
	biodata := []Biodata{
		{
			name:    "Rafli Ferdian",
			address: "Kendal",
			job:     "Mahasiswa",
			reason:  "Karena ingin mempelajari bahasa pemrograman baru terutama untuk Back-End",
		}, {
			name:    "John Doe",
			address: "Cityville",
			job:     "Insinyur Perangkat Lunak",
			reason:  "Tertarik untuk menjelajahi teknologi baru dan meningkatkan keterampilan pemrograman.",
		},
		{
			name:    "Jane Smith",
			address: "Townsville",
			job:     "Ilmuwan Data",
			reason:  "Penuh semangat tentang analisis data dan pembelajaran mesin.",
		},
		{
			name:    "Alice Johnson",
			address: "Village Heights",
			job:     "Pengembang Web",
			reason:  "Antusias membangun situs web interaktif dan ramah pengguna.",
		},
		{
			name:    "Bob Anderson",
			address: "Suburbia",
			job:     "Desainer Grafis",
			reason:  "Kreativitas dan desain adalah passion saya.",
		},
		{
			name:    "Eva Williams",
			address: "Metropolis",
			job:     "Manajer Produk",
			reason:  "Memimpin tim untuk memberikan produk yang sukses dan berdampak.",
		},
		{
			name:    "Michael Brown",
			address: "Rural Town",
			job:     "Administrator Sistem",
			reason:  "Memastikan operasional yang lancar dari sistem TI.",
		},
		{
			name:    "Sophia Davis",
			address: "Coastal City",
			job:     "Desainer UX/UI",
			reason:  "Menciptakan pengalaman pengguna yang mulus dan menarik secara visual.",
		},
		{
			name:    "Chris Miller",
			address: "Mountain Village",
			job:     "Insinyur Jaringan",
			reason:  "Membangun dan mengoptimalkan infrastruktur jaringan.",
		},
		{
			name:    "Olivia Taylor",
			address: "Island Paradise",
			job:     "Pengembang Game",
			reason:  "Mengubah ide menjadi pengalaman bermain game yang imersif.",
		},
	}

	// mengambil argumen dari terminal dan merubah ke int
	args, err := strconv.Atoi(argsRaw[1])
	if err != nil {
		fmt.Println("Error converting argument to integer:", err)
		return
	}
	fmt.Printf("Absen ke-%d\n", args)

	// memanggil function getBiodata
	getBiodata(biodata, args)
}

func getBiodata(data []Biodata, index int) {
	if index <= 0 || index >= len(data) {
		fmt.Println("Index out of range")
		return
	}

	biodata := data[index-1]

	// mencetak biodata
	fmt.Printf("Nama: %s\n", biodata.name)
	fmt.Printf("Alamat: %s\n", biodata.address)
	fmt.Printf("Pekerjaan: %s\n", biodata.job)
	fmt.Printf("Alasan memilih kelas Golang: %s\n", biodata.reason)
}
