package main

import "fmt"

func luas(s int)  {
	var hasil int
	hasil=s*s
	fmt.Println("LUAS:", hasil)
	
}
func keliling(s int)  {
	var hasil int
	hasil=s*4
	fmt.Println("KELILING:", hasil)
}

func main(){
	var s int
	fmt.Print("SISI:")
	fmt.Scan(&s)
	luas(s)
	keliling(s)
	
}
