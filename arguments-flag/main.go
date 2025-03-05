package main

import (
	"flag"
	"fmt"
)

func main() {
	// Arguments adalah data argumen opsional yang disisipkan ketika eksekusi program.
	// Flag merupakan ekstensi dari program.

	// Penggunaan Arguments melalui variable os.Args
	// Arguments disisipkan saat eksekusi program menggunakan
	// go run main.go  banana potato "ice cream" atau go build arguments
	// NOTE: untuk argumen dua kata atau lebih harus diapit "" agar tidak dihitung sebagai 2 argumen
	// argsRaw := os.Args
	// fmt.Printf("-> %#v\n", argsRaw)
	// []string{"/tmp/go-build3763114559/b001/exe/main","banana", "9 func main() {│potato", "ice cream"}

	// args := argsRaw[1:1]
	// fmt.Printf("-> %#v\n", args)
	// []string{}

	// Flag mirip dengan argument namun dengan penulisan berbentu key-value
	// Flag yang nilainya tidak di set akan mengambalikan nilai default
	// Cara penulisan argumen menggunakan flag
	// go run main.go -name="John Wick" -age=28
	// return value dari flag.String adalah string pointer sehingga perlu di dereference
	// name := flag.String("name", "anonymous", "type your name")
	age := flag.Int64("age", 25, "type your age")

	// fmt.Printf("name\t: %s\n", *name)
	// fmt.Printf("age\t: %d\n", *age)

	// Deklarasi flag dengan cara passing reference variable penampung data
	// mirip dengan cara sebelumnya namun nilainya diambil lewat parameter pointer

	// cara 1
	data1 := flag.String("name", "anonymous", "type your name")
	// fmt.Println(*data1)

	// cara 2
	var data2 string
	flag.StringVar(&data2, "gender", "male", "type your gender")
	flag.Parse()
	fmt.Printf("age\t: %d\n", *age)
	fmt.Println(*data1)
	fmt.Println(data2)
}
