package main

import(
  "fmt"
)

func main(){
			// penggunaan fallthrough dalam switch
			var point = 6

			switch {
			case point == 8:
				fmt.Println("perfect")
			case (point < 8) && (point > 3):
				fmt.Println("awesome")
				fallthrough
			case point < 5:
				fmt.Println("you need to learn more")
			default:
				{
					fmt.Println("not bad")
					fmt.Println("you can be better!")
				}
		  }
}
