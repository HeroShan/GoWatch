package prime
import(
	"testing"
	"fmt"
)
func TestIsprime(t *testing.T){
		for i:=0;i<1000;i++{
			p := Isprime(i)
			if p == true{
				fmt.Println("这个数字是质数：--",i)
			}
		}
		
}