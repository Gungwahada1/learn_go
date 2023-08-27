package main

import (
	"fmt"
	"time"
	"learn_go/packages/data_structure"
	"learn_go/packages/function"
	"learn_go/packages/map"
	"learn_go/packages/error_handling"
	"learn_go/packages/server"
	"github.com/go-sql-driver/mysql"
)

func main() {
	var name string = "Gungwah"
	var old int = 21
	var height float64 = 175.5
	isStudent := true

	fmt.Println("Halo,")
	fmt.Print("Salam Kenal, ")
	fmt.Printf("nama Saya %s, saya berumur %d, tinggi saya %.1f, dan saat ini saya menjadi mahasiswa: %t\n", name, old, height, isStudent)

	var (
		guest    string
		guestOld int
	)

	fmt.Printf("Siapa kamu? ")
	fmt.Scanf("%s", &guest)

	fmt.Printf("Berapa umurmu: ")
	fmt.Scan(&guestOld)

	// Pengkondisian If else
	if guest == name {
		fmt.Println("Nama kita sama!")
	} else {
		fmt.Println("Nama kita berbeda!")
	}
	if guestOld <= 18 {
		fmt.Println("Kamu masih kecil!")
	} else {
		fmt.Println("Kamu sudah dewasa!")
	}

	// Pengkondision Switch
	switch isStudent {
	case true:
		fmt.Println("Gungwah masih perlu belajar!")
	case false:
		fmt.Println("Gungwah sudah bisa bekerja!")
	}

	// Perulangan For loop
	fmt.Println("Perulangan For loop")
	for i := 1; i < guestOld; i++ {
		fmt.Print(i)
		fmt.Println("")
	}

	// Perulangan While
	fmt.Println("Perulangan While loop")
	var i int = 0
	for i <= guestOld {
		fmt.Print(i)
		i++
		fmt.Println("")
	}

	// Perulangan dengan Range
	fmt.Println("Perulangan dengan Range")
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	functions.TestingFunction(15, 10)

	data_structure.Slice()

	go func_map.Map()

	time.Sleep(5 * time.Second)

	result, err := error_handling.Divide(0,0)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	config := mysql.NewConfig()
	fmt.Println("MySQL Config:", config)

	server.TestingServer()
	server.TestingServeMux()
	server.HandlerTesting()
}
