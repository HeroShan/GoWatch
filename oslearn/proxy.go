package oslearn
// 代理模式的定义：由于某些原因需要给某对象提供一个代理以控制对该对象的访问。
// 访问对象不适合或者不能直接引用目标对象，代理对象作为访问对象和目标对象之间的中介。
type Subject interface {
	Do() string
}

type RealSubject struct{}

func (RealSubject) Do() string {
	return "real"
}

type Proxy struct {
	real RealSubject
}

func (p Proxy) Do() string {
	var res string

	// 在调用真实对象之前的工作，检查缓存，判断权限，实例化真实对象等。。
	res += "pre:"

	// 调用真实对象
	res += p.real.Do()

	// 调用之后的操作，如缓存结果，对结果进行处理等。。
	res += ":after"

	return res
}