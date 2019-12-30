package qa



type Astruct struct{
	Struct_id	int 	
	Auther		string
	Qtitle		string
	Qdescribe	string
	Content_id	string
	Date		int
	Issue		int
	Counts		int
}

type Element	struct{
	Qe_id		int
	Name		string
	Field		string
	Extra		string
}

type Content	struct{
	Content_id	int
	Question	string
	Anwser		string
	Qe_id		int
	Sort_id		int
	Required	int
}

type Anwser	struct{
	Anwser_id		int
	Content_id		int
	Anwser_content	string
	Struct_id		int
	User_ip			string
	Create_time		string
}

type ContentAnwser struct{
	Content `xorm:"extends"`
    Anwser 	`xorm:"extends"`
}

type User struct{
	Id 			int
	Username	string
	Password	string
	Lastime		int
}