package seckill
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
		//fmt.Printf("\r%02d:%02d:%02d",h,m,s)	
		fmt.Printf(" %c[%d;%d;%dm%s(f=%d,b=%d,d=%d)%c[0m ", 0x1B, 0, 47, 36, "", 36, 47, 0, 0x1B)
		os.Stdout.Sync()
	}
	
	
}
func split(r rune)bool{
	return r ==':'
}


