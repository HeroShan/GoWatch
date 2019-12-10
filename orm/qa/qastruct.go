package qa

import(
	"github.com/jinzhu/gorm"
)

type Astruct struct{
	gorm.Model
	struct_id	int 	`gorm:"primary_key"`
	auther		string
	title		string
	describe	string
	content_id	string
	date		int
	issue		int
	counts		int
}

type Element	struct{
	gorm.Model
	qe_id		int
	name		string
	field		string
	extra		string
}

type Content	struct{
	gorm.Model
	content_id	int
	question	string
	anwser		string
	qe_id		int
	sort_id		int
	required	int
}

type Anwser	struct{
	gorm.Model
	anwser_id		int
	content_id		int
	anwser_content	string
	struct_id		int
	user_ip			int
	create_time		string
}