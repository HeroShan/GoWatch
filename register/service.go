package register

type Service struct{
	Name	string	`json:"name"`
	Node	[]*Node	`json:"nodes"`
}

type Node struct{
	Ip		string	`json:"ip"`
	Port	int		`json:"port"`
	Id		int		`json:"id"`
	Weight	int		`json:"weight"`
}