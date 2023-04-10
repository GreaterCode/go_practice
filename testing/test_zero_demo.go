package main

import "fmt"

func main() {
	//for x := 1; x < 10; x++ {
	//	for y := 1; y <= x; y++ {
	//		fmt.Printf("%d*%d=%d ", x, y, x*y)
	//	}
	//	fmt.Println()
	//}
	OuterLoop:
		for i := 0; i < 2; i++ {
			for j := 0; j < 5; j++ {
				switch j {
				case 2:
					fmt.Println(i, j)
					break OuterLoop
				case 3:
					fmt.Println(i, j)
					break OuterLoop
				}
			}
	}
}
