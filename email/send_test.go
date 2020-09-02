package email

import(
	"testing"
	"log"
)

func TestSend(t *testing.T){
	//定义收件人
	mailTo := []string {
	"177868512@qq.com",
	"865134179@qq.com",
	}
  //邮件主题为"Hello"
   subject := "TEXT"
  // 邮件正文
   body := "Good"
  
	err := SendMail(mailTo, subject, body)
   log.Println(err)
}