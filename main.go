package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
    fmt.Print("Введите целое число: ") // ← Ошибка: не обновили текст
    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        input := scanner.Text()
        fmt.Printf("Вы ввели число: %s\n", input) // ← Ошибка: не обновили текст
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
