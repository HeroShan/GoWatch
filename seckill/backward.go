package main
import(
	"fmt"
	"time"
	"strings"
	"os"
	"strconv"
)
func backtime(){
	millisecond := time.Now().Unix() 
	now :=time.Unix(millisecond,0).Format("15:04:05")
	nowarr := strings.FieldsFunc(now,split)
	
	
	h,_ := strconv.Atoi(nowarr[0]) 
	m,_ := strconv.Atoi(nowarr[1]) 
	s,_ := strconv.Atoi(nowarr[2]) 
	
	for {
		s--
		if s == 0{
			s=60
			m--
			if m == 0{
				m=60
				h--
			}
		}
		time.Sleep(1 * time.Second)
		if s<10 {
			fmt.Printf("\r%d:%d:0%d",h,m,s)	
		}else if m<10 {
			fmt.Printf("\r%d:0%d:%d",h,m,s)
		}else if h<10 {
			fmt.Printf("\r0%d:%d:%d",h,m,s)
		}else{
			fmt.Printf("\r%d:%d:%d",h,m,s)
		}
		os.Stdout.Sync()
	}
	
	
}
func split(r rune)bool{
	return r ==':'
}
func main(){
	backtime()
}

