package main

import(
  "fmt"
)

func main(){
  var nama, matkul string
  var nilai, nim int

  fmt.Print("Masukkan NIM anda : ")
  fmt.Scan(&nim)

  fmt. Printf("Masukkan Nama anda : ")
  fmt.Scanf("%s", &nama)

  fmt. Printf("Masukkan Mata Kuliah anda : ")
  fmt.Scanf("%s", &matkul)

  fmt. Printf("Masukkan Nilai anda : ")
  fmt.Scanln(&nilai)

  fmt.Printf("Nama (nim) => %s %d \n", nama, nim)
  fmt.Printf("Matkul (nilai) => %s %d\n", matkul, nilai)

}
