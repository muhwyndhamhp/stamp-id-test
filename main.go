package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	opw "stamps-id/openweather"
	seq "stamps-id/uniquesequence"
	"strings"
)

func main() {

	fmt.Println("Selamat datang di program hasil test untuk Stamps Indonesia")
	fmt.Println("Silahkan pilih program yang akan dijalankan (cukup masukkan angka)")
	for {
		fmt.Println("1) FooBar bilangan prima")
		fmt.Println("2) Cuaca 5 hari kedepan")

		reader := bufio.NewReader(os.Stdin)
		program, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
			return
		}

		if strings.TrimSpace(program) == "1" {
			fmt.Println("\n====================")
			seq.UniquePrintSequence()
			fmt.Println("\n====================")
			break
		} else if strings.TrimSpace(program) == "2" {
			fmt.Println("\n====================")
			opw.RunForecast()
			fmt.Println("\n====================")
			break
		} else {
			fmt.Println("Program tidak dikenali, silahkan coba pilih angka kembali")
		}
	}

}
