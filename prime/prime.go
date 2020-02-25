package prime

import(
	_ "fmt"
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
	gcc := make(chan int,8)
	for i := 0; i <= N; i++{
		gcc <- i
		go func(i int){
			if Isprime(i) {
				nsclice = append(nsclice,i)
			}
			<-gcc
		}(i)
	}
	
	return nsclice
	
}