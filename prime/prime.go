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
		if P % i == 0 && P == i{
			break
		}else{
			if P % i == 0 || i>= P{
				return false
			}
			i++
		}
	}
	return true
}