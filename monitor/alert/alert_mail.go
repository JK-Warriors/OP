package alert

import (
	"sync"
	"fmt"
	"log"
    "net"
    "net/smtp"
	"crypto/tls"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	//"github.com/astaxie/beego/utils"
)

func AlertEMail(wg *sync.WaitGroup, mysql *xorm.Engine, alert_id int, retries int, sendto string, subject string, created int) {
	//添加异常处理
	defer func() {
		if err := recover(); err != nil{
		   // 出现异常，继续
		   log.Printf("Error: %v", err)
		   (*wg).Done()
		}
	}()
	
	//check golbal config
	sendmail_global := GetSendMailGlobal(mysql)
	if sendmail_global != "1" {
		log.Println("The global config for send mail is off.")
		(*wg).Done()
		return
	}

	
	if retries < 3 {
		host := GetSMTPHost(mysql)
		port := GetSMTPPort(mysql)
		user := GetSMTPUser(mysql)
		passwd := GetSMTPPassword(mysql)
		sendfrom := GetSendFrom(mysql)
		content := GetMailContent(mysql, alert_id)

		//sendto = "29036548@qq.com"
		if sendto == ""{
			sendto = sendfrom
		}
		
		auth := smtp.PlainAuth(
			"",
			user,
			passwd,
			host,
		)

		header := make(map[string]string)
		header["From"] = sendfrom
		header["To"] = sendto
		header["Subject"] = subject
		header["Content-Type"] = "text/html; charset=UTF-8"

		message := ""
		for k, v := range header {
			message += fmt.Sprintf("%s: %s\r\n", k, v)
		}
		message += "\r\n" + content

		
		sendto_list := strings.Split(sendto, ";")
	  
		err := SendMailUsingTLS(
			fmt.Sprintf("%s:%s", host, port),
			auth,
			user,
			sendto_list,
			[]byte(message),
		)

		if err != nil {
			log.Printf("AlertEMail failed: %s", err.Error())
			sql := `update pms_alerts set send_mail_status = 0, send_mail_retries = send_mail_retries + 1, send_mail_error = ? where id = ?`
			_, err = mysql.Exec(sql, err.Error(), alert_id)
		}else{
			log.Println("AlertEMail successful")
			sql := `update pms_alerts set send_mail_status = 1 where id = ?`
			_, err = mysql.Exec(sql, alert_id)
		}
		log.Println("AlertEMail end")

	}

	(*wg).Done()

}


//return a smtp client
func Dial(addr string) (*smtp.Client, error) {
    conn, err := tls.Dial("tcp", addr, nil)
    if err != nil {
        log.Println("Dialing Error:", err)
        return nil, err
    }
    //分解主机端口字符串
    host, _, _ := net.SplitHostPort(addr)
    return smtp.NewClient(conn, host)
}
  
//参考net/smtp的func SendMail()
//使用net.Dial连接tls(ssl)端口时,smtp.NewClient()会卡住且不提示err
//len(to)>1时,to[1]开始提示是密送
func SendMailUsingTLS(addr string, auth smtp.Auth, from string, to []string, msg []byte) (err error) {
    //create smtp client
    c, err := Dial(addr)
    if err != nil {
        log.Println("Create smpt client error:", err)
        return err
    }
    defer c.Close()
  
    if auth != nil {
        if ok, _ := c.Extension("AUTH"); ok {
            if err = c.Auth(auth); err != nil {
                log.Println("Error during AUTH", err)
                return err
            }
        }
    }
  
    if err = c.Mail(from); err != nil {
        return err
    }
  
    for _, addr := range to {
        if err = c.Rcpt(addr); err != nil {
            return err
        }
    }
  
    w, err := c.Data()
    if err != nil {
        return err
    }
  
    _, err = w.Write(msg)
    if err != nil {
        return err
    }
  
    err = w.Close()
    if err != nil {
        return err
    }
  
    return c.Quit()
}
