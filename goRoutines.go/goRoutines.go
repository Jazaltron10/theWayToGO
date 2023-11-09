package main

import (
	"fmt"
	// "math/rand"
	"sync"
	"time"
)
// creating a Mutex(Mutual Exclusion)
// var m = sync.Mutex{}

var m = sync.RWMutex{}
// These are simply counters for whenever you spawn a goroutine
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		// Use the go keyword to run a function concurrently
		wg.Add(1) //adding a waitgroup counter to ensure the program waits for the goroutine to finish before exiting
		go dbCall(i)
	}
	wg.Wait() //waits for the counter to go back down to 0, meaning that all the tasks have been completed, and the rest of the code will execute
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nThe results are %v", results)
}

func dbCall(i int) {
	//Simulate DB call delay
	// var delay float32 = rand.Float32() * 2000
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	// fmt.Println("\nThe result from the database is:", dbData[i])
	// m.Lock()
	// results = append(results, dbData[i])
	// m.Unlock()
	save(dbData[i])
	log()
	wg.Done() //This decrements the counter
}

/*
NOTE
Two processes writing to the same memory at the same time could lead to corrupt memory, this is why we use Mutexes 
it matters where you put your mutexes
We use the mutex to control writing to our slices to avoid corrupt memory.
 Drawback: the lock and unlock completely locks out other goroutines from accessing the results slice this is where the RWMutex comes in.
*/
func save (result string){
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log(){
	m.RLock()
	fmt.Printf("\nThe current results are: %v", results)
	m.RUnlock()

}