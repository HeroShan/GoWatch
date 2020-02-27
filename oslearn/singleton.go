package oslearn
//单例模式
// 如果once.Do(f)被多次调用，只有第一次调用会执行f，即使f每次调用Do 提供的f值不同。
// 需要给每个要执行仅一次的函数都建立一个Once类型的实例。
import(
	"sync"
)

type Singleton struct{
	Name	string 
	Age		int	   
}

var singleton *Singleton

var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}
