package main

import "fmt"

// Interface User
type User interface {
	CetakNama()
	GetInfo() string
	SetName(name string)
}

// Struct as Class
type Person struct {
	id      int
	name    string
	age     int
	address string
}

type Mahasiswa struct {
	Person // Embedding
}

// Method
// Receiver argument
func (p Person) GetInfo() string {
	return fmt.Sprintf("My name is %s, I'm %d years old.", p.name, p.age)
}

func (p Person) info() string {
	return p.GetInfo()
}

// Receiver argument (pointer)
func (p *Person) SetName(name string) {
	p.name = name
}

func (p Person) CetakNama() {
	fmt.Println(p.name)
}

// Fungsi Constructor
func NewPerson(name, address string, id, age int) User {
	return &Person{
		id:      id,
		name:    name,
		age:     age,
		address: address,
	}
}

func main() {
	// Instance class ke object
	// var person User = &Person{
	// 	Name:    "John",
	// 	Age:     20,
	// 	Address: "Jl. A",
	// }

	// Instance class ke object
	person := NewPerson("John", "Jl. A", 1, 20)

	// Memanggil method
	person.SetName("Adi")
	fmt.Println(person.GetInfo())
	person.CetakNama()

	mhs := Mahasiswa{}
	mhs.SetName("Bambang")
	mhs.age = 12
	fmt.Println("dipanggil dari child object", mhs.GetInfo())
}
