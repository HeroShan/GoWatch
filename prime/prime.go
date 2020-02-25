package prime

import(
	"fmt"
	"time"
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
		ch chan int
		quit chan bool
	)
	for i := 0; i <= (N/2-1); i++{
		if Isprime(i) {
			nsclice1 = append(nsclice1,i)
		}
		if i == (N/2-1) {
			ch <- 7
		}
	}
	go func(N2,N int){
		for j := N2; j <= N; j++{
			if Isprime(j) {
				nsclice2 = append(nsclice2,j)
			}
			if j == N {
				fmt.Println(nsclice2)
				ch <- 11
			}
		}
	}(N/2,N)
	for{
		select{
		case Pch := <-ch :
			if Pch == 7 {
				nsclice = append(nsclice,nsclice1...)
			}
			if Pch == 8 {
				nsclice = append(nsclice,nsclice2...)
			}
		case <-time.After(3 * time.Second) :
			quit<-true
		}

	}
	
	<-quit
	return nsclice
	
}