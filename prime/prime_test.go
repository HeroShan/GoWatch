package prime
import(
	"testing"
	"fmt"
)
func TestIsprime(t *testing.T){
	for i:=0;i<50;i++{
		p := Isprime(i)
		fmt.Println(p,",",i)
	}
}