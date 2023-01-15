package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(topKFrequent([]int{1, 2, 3, 1, 5, 8, 5}, 2))
}

//1.18
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	var queues []queue
	for _, v := range nums {
		m[v]++
	}
	for i, v := range m {
		queues = append(queues, queue{v, i})
	}
	sort.Slice(queues, func(i, j int) bool {
		return queues[i].n > queues[j].n
	})
	res := make([]int, k)
	for i, v := range queues[:k] {
		res[i] = v.p
	}
	return res
}

type queue struct {
	n int
	p int
}
