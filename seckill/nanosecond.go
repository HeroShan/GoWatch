package seckill

import (
	"time"
	"fmt"
)
func Fortime() string{
	
		millisecond := time.Now().Unix() 
		now :=time.Unix(millisecond,0).Format("15:04:05")
		return now
	
	
}
func OutTime(){
	i:=1
	for{
		sn := Fortime()
		if i==99{
			i=1
		}
		i++
		if i<10{
			fmt.Printf("\r%v[0%v]",sn,i)	
		}else{
			fmt.Printf("\r%v[%v]",sn,i)
		}
	}
}