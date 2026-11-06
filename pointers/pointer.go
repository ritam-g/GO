package main

import "fmt"

func normalFun(n int ) {
	n++
	fmt.Println(n)
}

func pointerFun(n *int){
	*n++
	fmt.Println(*n)
}

func main() {
	num := 0
	//normal fun which will not cahgne the the varible valaue for all it teat params is is varibal enot addrese 
	normalFun(num)
	fmt.Println(num)

	// now the thing for the point ther it will take the addrese of the varin so insted of treting parms as varible it will take address so wahter ever wil lchagne we can see global level got it 

	pointerFun(&num)
	fmt.Println(num)



}