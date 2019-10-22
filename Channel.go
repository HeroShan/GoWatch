package main
import (
	"time"
	"fmt"
)

var tmp int
func big(ch chan int){
	i := <-ch
	j := &tmp
	tmp = (*j+11)*i
	// fmt.Println(i,tmp)
	// fmt.Printf("%T",*j)
	// fmt.Println()
	// return j
}
func isPrime(value int) bool {
    if value <= 3 {
        return false
    }
    if value%2 == 0 || value%3 == 0 {
        return false
    }
    for i := 5; i*i <= value; i += 6 {
        if value%i == 0 || value%(i+2) == 0 {
            return false
        }
    }
    return true
}

func main(){
	sday := time.Now().Day()
	now := time.Now()
	ch := make(chan int,1)
    list  := make (map[int]time.Time)
	for i:=0;i<5;i++{
		bo:=isPrime(i)
		if i%2==1{
			ch<-i
			big(ch)
			list[sday+tmp] = now
		}else if bo == true{
			list[sday+i] = now.AddDate(-1,0,1)
		}else{
			list[sday+i] = now
		}
		 
		
	}
	
	for k,v := range list{
		fmt.Println(v,k)
		
	}


}