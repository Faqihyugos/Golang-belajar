package main

import(
  "fmt"
)

// var variable_name [ukuran1][ukuran2]...[ukuranN] type_variable
// ukuran1 menujukkan total array, sedangkan ukuran2 menunjukan jumlah data di dalam array ukuran1.
func main(){
  fmt.Println("Array Multi Dimensi")
  var angka = [5][2]int{{0,0}, {1,2}, {3,4}, {5,6}, {7,8}}

  for i := 0; i < 5; i++{
    for j := 0; j < 2; j++{
      fmt.Printf("Data ke [%d][%d] = %d\n", i, j, angka[i][j])
    }
  }

}



