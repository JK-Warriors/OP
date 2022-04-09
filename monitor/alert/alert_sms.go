package alert

import (
	"sync"
	"log"
	"github.com/xormplus/xorm"
)


func AlertSMS(wg *sync.WaitGroup, mysql *xorm.Engine, alert_id int, retries int, sendto string, subject string, created int) {
	//添加异常处理
	defer func() {
		if err := recover(); err != nil{
			// 出现异常，继续
			log.Printf("Error: %v", err)
			(*wg).Done()
		}
	}()

	

}
