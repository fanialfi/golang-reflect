package main

import (
	"fmt"
	"reflect"
)

func main() {
	persons := []map[string]any{
		{"nama": "fani", "umur": 17, "alamat": "indonesia"},
		{"nama": "alfi", "umur": 17, "alamat": "indonesia"},
		{"nama": "fanialfi", "umur": 17, "alamat": "indonesia"},
	}

	reflectValue1 := reflect.ValueOf(persons)

	if reflectValue1.Kind() == reflect.Slice {
		fmt.Printf("nilai variabel %v\nukuran variabel %d\n\n", reflectValue1.Slice(0, len(persons)), cap(persons))
	}

	fmt.Printf("tipe variabel :%v\n\n", reflectValue1.Type())
	for _, elm := range persons {
		fmt.Printf("tipe data elm : %v\n", (reflect.ValueOf(elm)).Kind()) // function Kind() mengembalikan tipe data nya
		fmt.Println(reflect.TypeOf(elm))                                  // mengembalikan struktur tipe data-nya
		fmt.Println(elm)
		fmt.Printf("nama saya %s, umur saya %d, dan alamat saya di %s\n\n", elm["nama"], elm["umur"], elm["alamat"])
	}

	// mencoba perbedaan antara function Kind() dengan Typeof()
	type person struct {
		nama string
		umur int
	}
	var interfaceKosong any = 10
	fmt.Println((reflect.ValueOf(interfaceKosong)).Kind(), reflect.TypeOf(interfaceKosong))

	interfaceKosong = []person{
		{nama: "fanialfi", umur: 17},
		{nama: "fani", umur: 17},
		{nama: "alfi", umur: 17},
	}
	fmt.Println((reflect.ValueOf(interfaceKosong)).Kind(), reflect.TypeOf(interfaceKosong))

	fmt.Println()

	for _, elm := range interfaceKosong.([]person) { // melakukan destructuring tipe data any ke bentuk asli-nya
		fmt.Println((reflect.ValueOf(elm)).Kind(), reflect.TypeOf(elm))
		fmt.Printf("nama saya %s, umur saya %d\n\n", elm.nama, elm.umur)
	}

	// pengaksesan nilai dalam bentuk interface{} / any
	fmt.Println()
	number := []person{
		{nama: "fani", umur: 17},
		{nama: "alfi", umur: 17},
		{nama: "fanialfi", umur: 17},
	}
	reflectValue2 := reflect.ValueOf(number)
	fmt.Println("tipe variabel", reflectValue2.Type())
	fmt.Println("nilai variabel", reflectValue2.Interface())

	// pengaksesan informasi properti variabel struct
	fmt.Println()
	s1 := &student{Name: "fani", Grade: 4}
	s1.getPropertiInfo()

	// pengaksesan informasi method pada struct
	fmt.Println()
	s1 = &student{Name: "fani", Grade: 4}
	fmt.Printf("nama\t: %s\n", s1.Name)

	reflectValue3 := reflect.ValueOf(s1)
	method := reflectValue3.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("fanialfi"),
	})

	fmt.Printf("nama\t: %s\n", s1.Name)
}

// pengaksesan informasi properti variabel struct

type student struct {
	Name  string
	Grade uint8
}

func (s *student) SetName(name string) {
	s.Name = name
}

func (s *student) getPropertiInfo() {
	reflectValue := reflect.ValueOf(s)

	// mengecek apakah variabel reflectvalue berupa pointer atau tidak
	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem() // mengambil isi aslinya dari pointer menggunakan Elm()
	}

	reflectType := reflectValue.Type()

	// function NumField() mengembalikan jumplah field public yang ada dalam struct
	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Printf("nama\t\t: %v\n", reflectType.Field(i).Name)          // mengembalikan nama field
		fmt.Printf("tipe data\t: %v\n", reflectType.Field(i).Type)       // mengembalikan tipe data field
		fmt.Printf("nilai\t\t: %v\n", reflectValue.Field(i).Interface()) // mengembalikan nilai dari field dalam bentuk interface{}
		fmt.Println("")
	}
}
