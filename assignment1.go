package main

import (
	"fmt"
	"os"
)

type Friend struct {
	id        uint
	name      string
	alamat    string
	pekerjaan string
	alasan    string
}

var friends []Friend = []Friend{
	{
		id:        1,
		name:      "Anto",
		alamat:    "Kuningan",
		pekerjaan: "Software Engineer",
		alasan:    "Karena golang laku dipasaran",
	},
	{
		id:        2,
		name:      "Yuni",
		alamat:    "Cipete",
		pekerjaan: "QA",
		alasan:    "Karena golang populer",
	},
	{
		id:        3,
		name:      "Aldo",
		alamat:    "Kebagusan Raya",
		pekerjaan: "Mahasiswa",
		alasan:    "Karena golang mudah dipelajari",
	},
	{
		id:        4,
		name:      "Abdul",
		alamat:    "Jl. Tebet Raya X",
		pekerjaan: "Mahasiswa",
		alasan:    "Karena golang memiliki banyak fitur",
	},
}

func main() {
	var arg = os.Args

	if len(arg) < 2 {
		fmt.Println("Tolong Masukkan Nama Teman")

	} else {
		findMyFriend(arg[1])
	}
}

func findMyFriend(name string) {
	found := false
	for _, eachFriend := range friends {
		if name == eachFriend.name {
			fmt.Println("ID:", eachFriend.id)
			fmt.Println("Nama:", eachFriend.name)
			fmt.Println("Alamat:", eachFriend.alamat)
			fmt.Println("Pekerjaan:", eachFriend.pekerjaan)
			fmt.Println("Alasan:", eachFriend.alasan)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Teman tidak ditemukan.")
	}
}
