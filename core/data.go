package core

import "sync"

var (
	instance = new(data)
	lock     = sync.Mutex{}
)

type data struct {
	m map[string]string
}

func GetSingleData() *data {
	once := sync.Once{}
	once.Do(func() {
		instance.m = make(map[string]string)
	})
	return instance
}

func (d *data) AddData(k string, v string) {
	lock.Lock()
	defer lock.Unlock()
	d.m[k] = v
}

func (d *data) GetData(k string) string {
	return d.m[k]
}
