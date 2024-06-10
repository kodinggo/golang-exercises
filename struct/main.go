package main

import "fmt"

type Mahasiswa struct {
	ID         int
	Nama       string
	IsVerified bool
	Jurusan    Jurusan
	Subjects   []Subject
}

type Jurusan struct {
	ID          int
	NamaJurusan string
}

type Subject struct {
	ID          int
	NamaSubject string
}

type Student struct {
	Mahasiswa
}

func main() {
	mhs := Mahasiswa{
		ID:         1,
		Nama:       "Bambang",
		IsVerified: true,
	}
	fmt.Println(mhs.ID, mhs.Nama, mhs.IsVerified)
	fmt.Print(mhs.Jurusan.NamaJurusan)

	// Literal struct
	person := struct {
		ID   int
		Nama string
	}{
		ID:   1,
		Nama: "Joko",
	}
	fmt.Println(person.ID, person.Nama)

	// Slice of struct
	mhss := []Mahasiswa{
		{
			ID:   1,
			Nama: "John",
		},
		{
			ID:   2,
			Nama: "David",
		},
	}
	for _, mhs := range mhss {
		fmt.Println(mhs.ID, mhs.Nama)
	}
}
