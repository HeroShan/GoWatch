package prime
import(
	"testing"
	"fmt"
)
func TestIsprime(t *testing.T){
		for i:=0;i<100000;i++{
			p := Isprime(i)
			if p == true{
				fmt.Println("这个数字是质数：--",i)
			}
		}
		
} 

func TestNprime(t *testing.T){
		n := Nprime(1231)
		fmt.Println(n)
}