package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func main() {

	jsonString, errJson := os.ReadFile("data.json")

	if errJson != nil {
		fmt.Println("Terjadi Kesalahan Saat Membaca JSON")
		return
	}

	jsonData := []byte(jsonString)

	type Phone struct {
		Vendor string `json:"vendor"`
		Model  string `json:"model"`
	}

	type VendorCount struct {
		Vendor string
		Total  int
	}

	//Soal 1
	fmt.Println("Soal 1 :")
	var data1 []map[string]string
	err1 := json.Unmarshal(jsonData, &data1)

	if err1 != nil {
		fmt.Println("Terjadi Kesalahan Saat Memasukkan Data JSON ke Map")
		return
	}

	for _, value1 := range data1 {
		if value1["vendor"] == "Samsung" {
			fmt.Println("Vendor : ", value1["vendor"], " | ", "Model: ", value1["model"])
		}
	}

	fmt.Println("\n")

	// Soal 2
	fmt.Println("Soal 2 : ")

	var Data2 []Phone

	err2 := json.Unmarshal(jsonData, &Data2)

	if err2 != nil {
		fmt.Println("Terjadi Kesalahan Saat Memasukkan Data JSON ke Struct")
		return
	}

	for _, value2 := range Data2 {
		if value2.Vendor == "Sony" {
			fmt.Println("Vendor : ", value2.Vendor, " | ", "Model: ", value2.Model)
		}
	}

	fmt.Println("\n")

	// Soal 3
	fmt.Println("Soal 3 : ")

	type AnimeCharacter struct {
		Name  string `json:"name"`
		Anime string `json:"anime"`
	}

	var Data3 []AnimeCharacter

	Data3 = append(Data3, AnimeCharacter{
		Name:  "Uchiha Sasuke",
		Anime: "Naruto",
	}, AnimeCharacter{
		Name:  "Yuta Okkotsu",
		Anime: "Jujutsu Kaiser",
	}, AnimeCharacter{
		Name:  "Ayanokoji Kyotaka",
		Anime: "Classroom of the Elite ",
	}, AnimeCharacter{
		Name:  "Levi Ackerman",
		Anime: "Attack On Titan",
	}, AnimeCharacter{
		Name:  "Johan Liebert",
		Anime: "Monster",
	})

	dataJson3, errMarshal3 := json.Marshal(Data3)

	if errMarshal3 != nil {
		fmt.Println("Terjadi Kesalahan Saat Menkonversi Struct Ke JSON")
		return
	}

	fmt.Println(string(dataJson3))

	fmt.Println("\n")

	//Soal 4
	fmt.Println("Soal 4 :")

	var Data4 []Phone

	err4 := json.Unmarshal(jsonData, &Data4)

	if err4 != nil {
		fmt.Println("Terjadi Kesalahan Saat Memasukkan Data JSON ke Struct", err4.Error())
		return
	}

	counts := make(map[string]int)

	for _, value4 := range Data4 {
		counts[value4.Vendor]++
	}

	var result []VendorCount

	for Vendor, Count := range counts {
		result = append(result, VendorCount{Vendor: Vendor, Total: Count})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Total > result[j].Total
	})

	fmt.Printf("%-20s | %s\n", "Vendor", "Jumlah Ponsel")
	fmt.Println("---------------------------------------")
	for _, res := range result {
		fmt.Printf("%-20s | %d\n", res.Vendor, res.Total)
	}

}
