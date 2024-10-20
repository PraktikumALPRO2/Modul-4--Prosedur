package main

import "fmt"

func faktorial(n int,hasil  *int) {
	*hasil= 1
	for i :=1; i<=n; i++{
		*hasil = *hasil *i
	}
}
func permutasi(n,r int, hasil *int) {
	var nfaktorial, nrfakrorial int
	faktorial(n, &nfaktorial)
	faktorial(r, &nrfakrorial)
	faktorial(n-r, &nrfakrorial)
	*hasil = nfaktorial/nrfakrorial
}
func kombinasi(n, r int, hasil *int){
	var nfaktorial, rfaktorial, nrfakrorial int
	faktorial(n, &nfaktorial)
	faktorial(r, &rfaktorial)
	faktorial(n-r, &nrfakrorial)
	*hasil = nfaktorial/(rfaktorial * nrfakrorial)
	
}
func main(){
	var a, b, c, d, p1, p2, c1, c2 int
	fmt.Print("Masukkan 4 bilangan: ")
	fmt.Scan(&a, &b, &c, &d)

	permutasi(a,c,&p1)
	kombinasi(a,c,&c1)
	fmt.Printf("%d %d\n",p1,c1)
	permutasi(b,d,&p2)
	kombinasi(b,d,&c2)
	fmt.Printf("%d %d\n",p2,c2)

}