package prime

import(
	"fmt"
)

func Isprime(P int) (bool) {
	var(
		i int = 2
	) 
	if P < 2 {
		return false
	}
	for {
		if i == P{
			fmt.Println("i=",i,"i=p")
			break
		}
		if P % i != 0{
			fmt.Println("i=",i,"P % i != 0")
			break
		}else{
			i++
			if i == P-1{
				fmt.Println("i=",i,"else P % i != 0")
				return false
			}
		}
	}
	return true
}