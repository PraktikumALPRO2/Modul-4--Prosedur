package main

import "fmt"

func main(){
	var a, b int
	var result int
	fmt.Scan(&a,&b)
	if  a >= b {
		permutasi(a,b,&result)
		fmt.Print(result)
	}else{
		permutasi(b,a,&result)
		fmt.Print(result)
	}
}

func faktorial(n int,hasil  *int) {
	*hasil= 1
	for i :=1; i<=n; i++{
		*hasil = *hasil *i
	}
}
func permutasi(n,r int, hasil *int) {
	var fakn, fakr int
	faktorial(n, &fakn)
	faktorial(n-r, &fakr)
	*hasil = fakn/fakr
}

