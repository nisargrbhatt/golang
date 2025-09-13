package main

import (
	"fmt"
	"sync"
	"time"
)

var dbData = []string{"id1", "id2", "id3", "id4", "id5", "id6", "id7", "id8", "id9", "id10"}

func dbCall(i int, results *[]string, lock *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	delay := 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	lock.Lock()
	*results = append(*results, dbData[i])
	lock.Unlock()

}

func main() {
	t0 := time.Now()

	// WaitGroup is used to wait for all goroutines to finish their execution
	var wg = sync.WaitGroup{}

	// Mutex or Mutual Exclusion is used to lock a resource so that only one goroutine can access it at a time else it will provide weird results
	var lock = sync.Mutex{}
	// Mutex blocks all operation to resource like read and write, so if you have need to block specific operation, use RWMutex. It has separate method for Read lock and Write lock
	// var rwlock = sync.RWMutex{}
	// rwlock.RLock();
	// defer rwlock.RUnlock();
	// rwlock.Lock()
	// defer rwlock.Unlock()
	var results = []string{}
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i, &results, &lock, &wg)
	}
	wg.Wait()
	fmt.Println("Results: ", results)
	fmt.Println("Time taken: ", time.Since(t0))
}
