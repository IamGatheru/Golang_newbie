package main

import (
    "fmt"
)

func main() {

    var troll string

    fmt.Println("Hello World folks!")
    fmt.Println("What's your name buddy?")
    fmt.Scan(&troll)

    fmt.Printf("Hello %v\n", troll)
}
