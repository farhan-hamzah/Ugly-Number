package main
import (
    "fmt"
)

// Fungsi untuk mengecek apakah suatu bilangan adalah ugly number
func isUgly(n int) bool {
    if n <= 0 {
        return false
    }

    for n%2 == 0 {
        n /= 2
    }
    for n%3 == 0 {
        n /= 3
    }
    for n%5 == 0 {
        n /= 5
    }

    return n == 1
}

func main() {
    var n int
    fmt.Print("Masukkan angka: ")
    fmt.Scanln(&n)

    if isUgly(n) {
        fmt.Println(n, "adalah ugly number")
    } else {
        fmt.Println(n, "bukan ugly number")
    }
}
