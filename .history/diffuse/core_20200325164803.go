package diffuse

import(
	_"github.com/gorilla/websocket"
	"time"
)

type Diffuse struct{
	Message		string;			//广播内容
	Sendtime	time.Time;		//发送时间
	Later		time.Second;	//延迟发送时间
	Title		string;			//广播标题
	Frequency	*Freq;			//

}

type Freq struct{
	Times		int;			//次数
	Interval	time.Second;	//间隔时间
}