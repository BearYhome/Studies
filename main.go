package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
    fmt.Print("Введите данные: ") //  Исправлено
    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        input := scanner.Text()
        fmt.Printf("Вы ввели следующие данные: %s\n", input) //  Исправлено
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
