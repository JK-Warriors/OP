package alert

import (
	"sync"
	"github.com/xormplus/xorm"
)


func AlertSMS(wg *sync.WaitGroup, mysql *xorm.Engine, alert_id int, retries int, sendto string, subject string, created int) {
}
