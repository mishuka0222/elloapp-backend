package sqlx

import (
	"time"
)

type ExecResult struct {
	LastInsertId int64
	RowsAffected int64
}

type StoreResult struct {
	Data interface{}
	Cost int
	Err  error
}

type StoreChannel chan StoreResult

func Do(f func(result *StoreResult)) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		result := StoreResult{}
		f(&result)
		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel
}

func Must(sc StoreChannel) interface{} {
	r := <-sc
	if r.Err != nil {
		time.Sleep(time.Second)
		panic(r.Err)
	}

	return r.Data
}
