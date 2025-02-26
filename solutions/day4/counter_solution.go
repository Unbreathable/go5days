package main

/*
// ID -> *Counter
var counterMap *sync.Map = &sync.Map{}

// Get a counter by its id or create a new one starting at 0
func getCounter(id string) *Counter {
	obj, _ := counterMap.LoadOrStore(id, &Counter{
		mutex: &sync.Mutex{},
		count: 0,
	})
	return obj.(*Counter)
}

type Counter struct {
	mutex *sync.Mutex
	count int64 // Never read this directly, should be protected using a mutex
}

// Increment the counter by one
func (c *Counter) Increment() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.count++
}

func (c *Counter) GetValue() int64 {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.count
}
*/
