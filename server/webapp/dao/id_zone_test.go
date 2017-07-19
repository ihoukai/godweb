package dao

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	. "gopkg.in/check.v1"
	"sync"
)

// TestGenerateID 并发测试生成id
func (s *TestSuite) TestGenerateID(c *C) {
	num := 100
	curID := _curIDZone.CurID
	wg := new(sync.WaitGroup)
	// 并发测试id
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			GenerateAccountID()
			wg.Done()
		}()
	}
	wg.Wait()
	if curID+int64(num) != _curIDZone.CurID {
		c.Errorf("Not Equal, Expect=%d Cur=%d", curID+int64(num), _curIDZone.CurID)
	}
}
