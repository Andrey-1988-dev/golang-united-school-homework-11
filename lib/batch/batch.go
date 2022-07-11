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
	var ch = make(chan int, pool)
	var wg sync.WaitGroup
	var mu sync.Mutex
	var i int64
	for i = 0; i < n; i++ {
		wg.Add(1)
		ch <- 1
		go func(j int64) {
			defer wg.Done()
			defer mu.Unlock()
			itemUser := getOne(j)
			mu.Lock()
			res = append(res, itemUser)
			<-ch
		}(i)
	}
	wg.Wait()
	return res
}
