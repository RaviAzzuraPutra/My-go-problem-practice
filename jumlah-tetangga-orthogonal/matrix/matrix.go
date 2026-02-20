package matrix

func PerhitunganMatrix(matrix [][]int, longitude int, latitude int) int {

	hasil := 0

	Baris := len(matrix)
	Kolom := len(matrix[0])

	if longitude-1 >= 0 {
		hasil += matrix[longitude-1][latitude]
	}

	if longitude+1 < Baris {
		hasil += matrix[longitude+1][latitude]
	}

	if latitude-1 >= 0 {
		hasil += matrix[longitude][latitude-1]
	}

	if latitude < Kolom {
		hasil += matrix[longitude][latitude+1]
	}

	return hasil

}
