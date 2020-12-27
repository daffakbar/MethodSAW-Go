package main

import (
	"fmt"
)

// fungsi untuk mendapatkan kolom array
func arrayColumn(board [3][4]int, columnIndex int) (column []int) {
	column = make([]int, 0)
	for _, row := range board {
		column = append(column, row[columnIndex])
	}
	return
}

// fungsi untuk mendapatkan minimal
var getMin = func(n []int) int {
	var min int
	for i, e := range n {
		switch {
		case i == 0:
			min = e
		// case e > max:
		// 	max = e
		case e < min:
			min = e
		}
	}
	return min
}

// fungsi untuk mendapatkan maximal
var getMax = func(n []int) int {
	var max int
	for i, e := range n {
		switch {
		case i == 0:
			max = e
		case e > max:
			max = e
			// case e < min:
			// 	min = e
		}
	}
	return max
}

func main() {

	// Menentukan rekomendasi tempat bootcamp IT
	// C1 = Biaya (seberapa banyak biaya untuk bootcamp)
	// C2 = Lokasi (seberapa jauh lokasi tempat bootcamp dari rumah)
	// C3 = Fasilitas ( fasilitas tempat bootcamp)
	// C4 = Kualitas pengajar (dilihat dari gelar akademik para mentor di tempat bootcamp)

	kriteria := [4][1]string{{"C1"}, {"C2"}, {"C3"}, {"C4"}}
	// Pembobotan nilai alternatif
	alternatif := [3][4]int{
		{3, 2, 3, 4}, //Alternatif 1 (Bootcamp Arkademy)
		{2, 4, 3, 3}, //Alternatif 2 (Bootcamp Hactiv8)
		{4, 5, 5, 5}} //Alternatif 3 (Bootcamp Dumbways)

	fmt.Println(kriteria)
	// fmt.Println(alternatif)
	// fmt.Println(alternatif[2])
	// das := arrayColumn(alternatif, 0)
	// // fmt.Println(ArrayColumn(alternatif {2})
	// fmt.Println(das)

	// var min = getMin(das)
	// println(min)
	// var max = getMax(das)
	// println(max)
	rr := [3][4]float32{}

	indexAlternatif := 0

	for _, element1 := range alternatif {
		indexKriteria := 0
		fmt.Println(element1)
		for _, element2 := range kriteria {
			// C1 dan C2 dicari nilai terkecil karena COST
			if string(element2[0]) == "C1" || element2[0] == "C2" {
				rr[indexAlternatif][indexKriteria] = float32(getMin(arrayColumn(alternatif, indexKriteria))) / float32(alternatif[indexAlternatif][indexKriteria])
				// C3 dan C4 dicari nilai terbesar karena BENEFIT
			} else if element2[0] == "C3" || element2[0] == "C4" {
				rr[indexAlternatif][indexKriteria] = float32(alternatif[indexAlternatif][indexKriteria]) / float32(getMax(arrayColumn(alternatif, indexKriteria)))
			}
			// fmt.Println(element)
			indexKriteria++
		}
		indexAlternatif++
	}
	fmt.Println("===Hasil normalisasi===")

	fmt.Println(rr)

	// Bobot kriteria C1, C2, C3, C4
	w := [4]float32{0.4, 0.3, 0.2, 0.1}
	fmt.Println(w)
	nv := [3]float32{}
	indexAlternatif1 := 0
	const indexV = 0
	fmt.Println("=================")

	// Bobot kriteria dikalikan hasil normalisasi
	for _, element := range alternatif {
		indexw := 0
		indexr := 0
		v := float32(0)
		// GAK PENGEN NYETAK IKI
		fmt.Println(element)
		// IKI
		for _, element2 := range kriteria {
			v += float32(w[indexw] * rr[indexAlternatif1][indexr])
			indexr++
			indexw++
			// GAK PENGEN NYETAK IKI
			fmt.Println(element2)
			// GAK PENGEN NYETAK IKI

		}
		// println("=======df============")

		nv[indexAlternatif1] = v
		indexAlternatif1++
		// fmt.Println(nv)
	}
	println("==Hasil perhitungan==")
	fmt.Println(nv)
	println("======================")

}
