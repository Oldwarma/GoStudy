package main

import (
	"fmt"
	"sync"
)

func main() {
	NewCache().add(st, &cache{
		Name:    "鸡哥",
		SurName: "cxk",
		Id:      "5178",
	})
	NewCache().print()
	NewCache().look("niganma")
	NewCache().look("你干嘛")
	NewCache().delete("你干嘛")
	NewCache().look("你干嘛")
}

type inf interface {
	add(s string, ca *cache)
	delete(s string)
	look(s string) *cache
}

func NewCache() *caches {
	return &caches{}
}

var (
	m  = make(map[string]*cache, 1)
	ch = make(map[string]chan interface{}, 1)
	st = "你干嘛"
)

type caches struct {
	cache
	mu sync.Mutex
}
type cache struct {
	Name    string
	SurName string
	Id      string
}

func (c *caches) add(s string, ca *cache) {
	c.mu.Lock()
	defer c.mu.Unlock()
	m[s] = ca
	fmt.Println("添加成功")
}
func (c *caches) delete(s string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(m, s)
}

func (c *caches) look(s string) *cache {
	_, ok := m[s]
	if ok {
		n := m[s]
		fmt.Println("该用户", n)
		return n
	} else {
		fmt.Println("无该用户", s)
	}
	return nil
}
func (c *caches) print() {
	for _, v := range m {
		fmt.Println("print", v)
	}

}
