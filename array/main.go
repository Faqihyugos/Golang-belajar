package main

import(
  "fmt"
)

func main(){
  var name [5]string

  name[0] = "faqih"
  name[1] = "yugo"
  name[3] = "yogi"
  name[4] = "alfaqih"

  for i, v := range name{
    fmt.Println("Index ke", i, "=>", v)
  }
}
