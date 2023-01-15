package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	tree := create(20)
	fmt.Println("tree value", tree.Value)
	tree = insert(tree, 10)
	tree = insert(tree, 30)
	traverse(tree)
	fmt.Println("tree value", tree.Value)
}

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

//递归访问二叉树上的所有节点
func traverse(t *Tree) {
	if t == nil {
		return
	}
	traverse(t.Left)
	fmt.Print(t.Value, " ")
	traverse(t.Right)
}

func create(n int) *Tree {
	var t *Tree
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		temp := rand.Intn(n)
		t = insert(t, temp)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v == t.Value {
		return t
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}
