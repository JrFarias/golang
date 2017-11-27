package main

import "fmt"
func olar(done chan string) {
    fmt.Println("2 Goroutine")
    done <- "A segunda voltou"
}
func olar2(result chan bool) {
    fmt.Println("3 Goroutine")
    result <- true
}
func hello(done chan string) {  
    fmt.Println("Hello world goroutine")
    result := make(chan bool)
    go olar(done)
    go olar2(result)
     <- result
}
func main() {  
    done := make(chan string)
    go hello(done)
    fmt.Println("olar aqui",<-done)
    fmt.Println("main function")
}