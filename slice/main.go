package main

import(
  "fmt"
)

func main(){
  var mahasiswa []string
  mahasiswa = []string{"faqih","susilo"}

  fmt.Println("Data mahasiswa ke 0 sampai 1 => ", mahasiswa[0:1])
  fmt.Println("Data mahasiswa ke 0 sampai 2 => ", mahasiswa[0:2])

  fmt.Println("Jumlah mahasiswa adalah => ", len(mahasiswa))

}
