package qa


type Astruct struct{
	struct_id	int 	
	auther		string
	qtitle		string
	qdescribe	string
	content_id	string
	date		int
	issue		int
	counts		int
}

type Element	struct{
	qe_id		int
	name		string
	field		string
	extra		string
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
	anwser_id		int
	content_id		int
	anwser_content	string
	struct_id		int
	user_ip			int
	create_time		string
}