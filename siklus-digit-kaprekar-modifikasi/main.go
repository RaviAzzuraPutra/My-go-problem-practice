package main

import (
	"fmt"
	"sort"
	"strconv"
)

func SiklusDigitKaprekarModifikasi(nums int) int {
	IntToStr := strconv.Itoa(nums)

	array1 := []int{}

	for _, value1 := range IntToStr {
		StrToInt, _ := strconv.Atoi(string(value1))

		array1 = append(array1, StrToInt)
	}

	asc := append([]int(nil), array1...)
	desc := append([]int(nil), array1...)

	sort.Ints(asc)

	sort.Sort(sort.Reverse(sort.IntSlice(desc)))

	var arrayToStrAsc string
	var arrayToStrDesc string

	for _, valueAsc := range asc {
		arrayToStrAsc += strconv.Itoa(valueAsc)
	}
	for _, valueDesc := range desc {
		arrayToStrDesc += strconv.Itoa(valueDesc)
	}

	StrToIntAsc, _ := strconv.Atoi(string(arrayToStrAsc))
	StrToIntDesc, _ := strconv.Atoi(string(arrayToStrDesc))

	result := StrToIntDesc - StrToIntAsc
	var iterasi int = 1

	for iterasi = 1; iterasi < 100; iterasi++ {
		if result == nums {
			return iterasi
		} else if result == 0 {
			return 0
		} else if result != nums {
			SiklusDigitKaprekarModifikasi(result)
		} else if iterasi > 100 {
			return 100
		}
	}

	return iterasi

}

func main() {
	Angka := 916

	fmt.Println(SiklusDigitKaprekarModifikasi(Angka))
}
