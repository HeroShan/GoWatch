package oslearn
//单例模式
import(
	"sync"
)

type Singleton struct{
	Name	string
	Age		int
}

var singleton *Singleton

var once sync.Once

func GetInstance(s Singleton) *Singleton {
	once.Do(func() {
		singleton = &s
	})
	return singleton
}
