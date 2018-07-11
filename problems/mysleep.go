package main

import (
  "fmt"
  "time"
)

func main() {
    fmt.Println("start sleep");
    mysleep(3)
    fmt.Println("end");
    c := make(chan int, 10)
    fmt.Println(c)
}

func mysleep(n int) {
    c := <- time.After(time.Second * time.Duration(n))
    fmt.Println(c)
}
