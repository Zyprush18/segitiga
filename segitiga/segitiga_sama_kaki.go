package segitiga
import "fmt"

func SegitigaSamaKaki(cetak string,tinggi int)  {
	height := tinggi

	for i := 1; i <= height; i++ {

		for j := 1; j <= height-i; j++ {
			fmt.Print(" ")
		}

		for k := 1; k <= 2*i-1; k++ {
			fmt.Print(cetak)
		}

		fmt.Println()

	}
}

func SegitigaSamaKakiTerbalik(cetak string,tinggi int)  {
	height := tinggi

	for i := height; i >= 1; i-- {

		for j := 1; j <= height-i; j++ {
			fmt.Print(" ")
		}

		for k := 1; k <= 2*i-1; k++ {
			fmt.Print(cetak)
		}

		fmt.Println()

	}
}

