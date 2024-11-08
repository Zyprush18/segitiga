package main

import (
	"fmt"

	"github.com/Zyprush18/segitiga/segitiga"
)

func main() {
	fmt.Println("\n --- Segitiga sama kaki ---")
	segitiga.SegitigaSamaKaki("*",5)
	fmt.Println("\n --- Segitiga sama kaki terbalik ---")
	segitiga.SegitigaSamaKakiTerbalik("*",5)
}