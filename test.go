package main

import "fmt"

// type student struct {
// 	// name  string
// 	grade int
// 	age   int
// 	person
// }
type student struct {
	person struct {
		name string
		age  int
	}
	grade int
	age   int
	hobi  []string
}

// type person struct {
// 	name string
// 	age  int
// }

func test() {
	// Nested struct
	var s1 = student{}
	s1.person.name = "daffa"
	s1.person.age = 42
	s1.grade = 2
	s1.age = 42
	fmt.Println("name", s1.person.name)
	fmt.Println("age", s1.person.age)
	fmt.Println("grade", s1.grade)

	// var siswa = struct {
	// 	person
	// 	grade int
	// }{}
	// siswa.person = person{"eman", 14}
	// siswa.grade = 2
	// fmt.Println("name", siswa.name)
	// fmt.Println("age", siswa.age)
	// fmt.Println("grade", siswa.grade)

	// struct
	// var siswa = student{"Daffa", 5}
	// var siswa = student{name: "Daffa"}
	// siswa.name = "daffa akbar"
	// siswa.grade = 2

	// var siswa = student{}
	// siswa.name = "eman"
	// siswa.grade = 21
	// siswa.age = 55
	// siswa.person.age = 445

	// fmt.Println("name", siswa.name)
	// fmt.Println("age", siswa.age)
	// fmt.Println("grade", siswa.grade)
	// fmt.Println("grade", siswa.person.age)
	// /////////////////////////////////////
	///////////////////////////////////////
	// Pagi tgl 15
	// fmt.Println("Siapa nama anda?")

	// var nama string
	// fmt.Scanln(&nama)
	// fmt.Println("Nama anda " + nama)

	// var nama string = "Dio"
	// fmt.Println(nama)

	// nama = "Akbar"
	// fmt.Println("nama Kedua " + nama)

	// const umur int8 = 23
	// fmt.Println(umur)

	// siang
	// var warna string = "merah"
	// if warna == "merah" {
	// 	fmt.Println("berhenti")
	// } else if warna == "kuning" {
	// 	fmt.Println("hati-hati")
	// } else {
	// 	fmt.Println("Jalan")
	// }
	// var warna string = "merah"

	// switch warna {
	// case "hijau":
	// 	fmt.Println("Jalan")
	// case "kuning":
	// 	fmt.Println("hati2")
	// default "merah":
	// 	fmt.Println("berhenti")
	// }

	// for i := 0; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// for i := 1; i <= 10; i++ {
	// 	if i&2 == 0 {
	// 		continue
	// 	}
	// 	if i > 8 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// Array
	// var name [3]string
	// name[0] = "Dio"
	// name[1] = "Daffa"
	// name[2] = "Akbar"

	// fmt.Println(name[0], name[1], name[2])

	// var name = [3]string{"Dio", "Daffa", "Akbar"}
	// fmt.Println(name[0], name[1], name[2])

	// var name = [...]string{"Dio", "Daffa", "Akbar"}
	// fmt.Println(len(name))
	// array bertumpuk
	// var data = [2][3]int{[3]{1,2,3}}, [3]int{3,4,5}

	// var data1 = [2][3]int{{1, 2, 3}, {3, 4, 5}}
	// fmt.Println(data1)

	// var name = [3]string{"Dio", "Daffa", "Akbar"}
	// for biasa
	// for i := 0; i < len(name); i++ {
	// 	fmt.Println("name ", name[i])
	// }
	// for range
	// for index, value := range name {
	// 	fmt.Println("nama ", value, "itu index ke ", index)
	// }

}
