package main

import(
  "fmt"
)

func main(){
  var index = 0

  for index < 10 {
    fmt.Println("Angka ke -",index)
    index++

    if index == 5 {
      break
    }

  }
}
