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
	ch := make(chan int, pool)
	var wg sync.WaitGroup
	var mu sync.Mutex
	var users []user
	for i := 0; i < int(n); i++ {
		wg.Add(1)
		ch <- 1
		go func(j int, u []user) {
			defer wg.Done()
			defer mu.Unlock()
			itemUser := getOne(int64(j))
			mu.Lock()
			users = append(users, itemUser)
			<-ch
		}(i, users)
	}
	wg.Wait()
	return users
}
