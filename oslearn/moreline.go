package oslearn

import(
	"os"
	"fmt"
	"time"
)
func DD(){
	fileName := "../css/img/px.png"
	desfileName := "../css/img/CCpx.png"
	sfile,err := os.Open(fileName)
	if err != nil{
		fmt.Println(nil)
	}

	info,_ := os.Stat(fileName)
	size := info.Size()
	var scount  int64 = 1
	if size % 2 == 0{
		scount *= 2
	}else if size % 3 == 0 {
		scount *= 3
	}else{
		scount *= 1
	}

	si := size / scount
	fmt.Printf("文件总大小：%v, 分片数：%v,每个分片大小：%v\n",size,scount,si)

	desF,err := os.OpenFile(desfileName,os.O_CREATE|os.O_RDWR,0755)
	if err != nil{
		fmt.Println(err)
	}

	for i:=0;i<int(si);i++{
		go func(vs int){
			//申明一个byte
			b := make([]byte,si)
			//从指定位置开始读
			sfile.ReadAt(b,int64(vs)*si)
			//从指定位置开始写
			desF.WriteAt(b,int64(vs)*si)

		}(i)
	}
	time.Sleep(time.Second*5)
	defer desF.Close()
	defer sfile.Close()
}