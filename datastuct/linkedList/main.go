package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

var root = new(Node)

func main() {
	fmt.Println(root)
	root = nil
	traverse(root)
	addNode(root, 1)
	addNode(root, 10)
	if lookupNode(root, 11) {
		fmt.Println("Node exists!")
	} else {
		fmt.Println("Node does not exist!")
	}
}

//添加节点
func addNode(t *Node, v int) int {
	if root == nil {
		t = &Node{v, nil}
		root = t
		return 0
	}
	if v == t.Value {
		fmt.Println("node already exists:", v)
		return -1
	}
	if t.Next == nil {
		t.Next = &Node{v, nil}
		return -2
	}
	return addNode(t.Next, v)
}

//递归访问链表
func traverse(t *Node) {
	if t == nil {
		fmt.Println("->Empty list!")
		return
	}
	for t != nil {
		fmt.Printf("%d->", t.Value)
		t = t.Next
	}
	fmt.Println()
}

//查看节点是否存在
func lookupNode(t *Node, v int) bool {
	if root == nil {
		t = &Node{v, nil}
		root = t
		return false
	}
	if v == t.Value {
		return true
	}
	if t.Next == nil {
		return false
	}
	return lookupNode(t.Next, v)
}

//长度
func size(t *Node) int {
	if t == nil {
		fmt.Println("-> empty list!")
		return 0
	}
	i := 0
	for t != nil {
		i++
		t = t.Next
	}
	return i
}
