package main

import (
	"fmt"
	"jumlah-tetangga-orthogonal/matrix"
)

func main() {

	ArrayMatrix := [][]int{
		{0, 9, 8},
		{7, 6, 5},
		{4, 3, 2},
	}
	longitude := 0
	latitude := 0

	fmt.Println(matrix.PerhitunganMatrix(ArrayMatrix, longitude, latitude))

}
