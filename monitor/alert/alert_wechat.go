package alert

import (
	"sync"
	"github.com/xormplus/xorm"
)



func AlertWeChat(wg *sync.WaitGroup, mysql *xorm.Engine, alert_id int, retries int, subject string, created int) {
}
