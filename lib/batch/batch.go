package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	var wg sync.WaitGroup
	var m sync.Mutex
	wg.Add(int(n))
	c := make(chan struct{}, pool)

	for i := int64(0); i < n; i++ {
		c <- struct{}{}
		go func(id int64) {
			defer wg.Done()
			user := getOne(id)
			m.Lock()
			res = append(res, user)
			m.Unlock()
			<-c
		}(i)

	}
	wg.Wait()

	return res
}
