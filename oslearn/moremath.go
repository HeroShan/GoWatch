package oslearn
import(
	"fmt"
	"time"
)
var num,i int
func Cc(){
	
	num = 3
	for i=1;i<10;i++{
		
		go func(i int){
			num = i^num<<num+11
			//fmt.Println(num)
		}(i)
		go func(i int) {
			num = i+num+7
			//fmt.Println(num)
		}(i)
		go func(i int) {
			
			num = i*13<<2
			//fmt.Println(num)
		}(i)
		time.Sleep(time.Nanosecond*1777)
	}
	defer fmt.Println(num)
	//fmt.Println(num)
}

