package diffuse

import(
	_"github.com/gorilla/websocket"
	"time"
)

type Diffuse struct{
	Message		string;			//广播内容
	Sendtime	time.Time;		//发送时间
	Later		time.Duration ;	//延迟发送时间
	Title		string;			//广播标题
	Frequency	*Freq;			//展示评率

}

type Freq struct{
	Times		int;			//次数
	Interval	time.Duration ;	//间隔时间
}

func Init(){
	 
}

func DiffuseMutex(){
	
}