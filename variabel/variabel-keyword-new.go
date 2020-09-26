package main

import "fmt"

func main() {
  		  // variabel dengan keyword new
        name := new(string)

        fmt.Println(name) // 0xc000010200
        fmt.Println(*name) // ""

}
