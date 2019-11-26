package heartbeat
import(
	"os"
	"strings"

)

type Alive struct{
	ip 		string
	status	*State
}

type State struct{
	on		string
	off		string
	timeout string
}

func GetConfIp()[]string{
	file,_ := os.Open("../fsm/fsm.config")
	data := make([]byte,512)
	var filestr string
	for{
		count, _ := file.Read(data)
		filestr += string(data[:count])
		if count == 0{
			break
		}
	}
	ip := strings.Split(filestr,"=")
	return strings.Split(ip[1],",")
}

func CheckStatus(clientIp string){

}



