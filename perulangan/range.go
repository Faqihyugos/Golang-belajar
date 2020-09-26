package main

import(
  "fmt"
)

func main(){
  data := []string{"Faqih","Yugo","Susilo"}

  for index, v := range data {
    fmt.Println("Perulangan ke", index, v)
  }
}
