package prime

import(
	_"fmt"
)

func Isprime(P int) (bool) {
	var(
		i int = 2
	) 
	if P < 2 {
		return false
	}
	for {
		if  P == i{
			break
		}else{
			if P % i == 0 {
				return false
			}
			i++
		}
	}
	return true
}

func Nprime(N int)(nsclice []int){
	var(
		nsclice1 []int
		nsclice2 []int
		ich chan int
		jch chan int
	)
	go func(N2,N int){
		for j := N2; j <= N; j++{
			if Isprime(j) {
				nsclice2 = append(nsclice2,j)
			}
			if j == N {
				jch <- 11
			}
		}
	}(N/2,N)
	go func(N int){
		for i := 0; i <= (N/2-1); i++{
			if Isprime(i) {
				nsclice1 = append(nsclice1,i)
			}
			if i == (N/2-1) {
				ich <- 7
			}
		}
	}(N/2)
	
	for{
		var ic,jc int
		select{
			case <-ich :
					nsclice = append(nsclice,nsclice1...)
					ic = 2
			case <-jch :
					nsclice = append(nsclice,nsclice2...)
					jc = 2
			if ic + jc == 4{
				goto LOOP
			}
		}

	}
	
	
	
	LOOP: return nsclice
	
}