package main

import "fmt"

// 定义节点

type DoubleNode struct {
	Value    int
	Previous *DoubleNode
	Next     *DoubleNode
}

// 添加节点
func addNode1(t *DoubleNode, v int) int {
	if DoubleHead == nil {
		t = &DoubleNode{v, nil, nil}
		DoubleHead = t
		return 0
	}

	if v == t.Value {
		fmt.Println("节点已经存在：", v)
		return -1
	}
	// 如果当前节点下一个节点为空
	if t.Next == nil {
		temp := t
		// 与单链表不同的是每个节点还要维护前驱节点指针
		t.Next = &DoubleNode{v, temp, nil}
		return -2
	}
	return addNode1(t.Next, v)
}

// 遍历链表

func traverse1(t *DoubleNode) {
	if t == nil {
		fmt.Println("-> 空链表")
		return
	}

	for t != nil {
		fmt.Printf("%d ->", t.Value)
		t = t.Next
	}
	fmt.Println()
}

// 反向遍历链表
func reverse(t *DoubleNode) {
	if t == nil {
		fmt.Println("->  空链表")
		return
	}

	temp := t
	for t != nil {
		temp = t
		t = t.Next
	}

	for temp.Previous != nil {
		fmt.Printf("%d -> ", temp.Value)
		temp = temp.Previous
	}
	fmt.Printf("%d -> ", temp.Value)
	fmt.Println()
}

// 获取链表长度
func Dsize(t *DoubleNode) int {

	if t == nil {
		fmt.Println("-> 空链表")
		return 0
	}
	n := 0
	for t != nil {
		n++
		t = t.Next

	}
	return n
}

//查找节点
func lookUPDNode(t *DoubleNode, v int) bool {
	if DoubleHead == nil {
		return false
	}

	if v == t.Value {
		return true
	}

	if t.Next == nil {
		return false
	}

	return lookUPDNode(t.Next, v)
}

var DoubleHead = new(DoubleNode)

func main() {
	fmt.Println(DoubleHead)
	DoubleHead = nil
	// 遍历链表
	traverse1(DoubleHead)
	// 新增节点
	addNode1(DoubleHead, 1)
	// 再次遍历
	traverse1(DoubleHead)
	// 继续添加节点
	addNode1(DoubleHead, 10)
	addNode1(DoubleHead, 5)
	addNode1(DoubleHead, 100)
	// 再次遍历
	traverse1(DoubleHead)

	// 添加已存在节点
	addNode1(DoubleHead, 100)
	fmt.Println("链表长度:", Dsize(DoubleHead))
	// 再次遍历
	traverse1(DoubleHead)
	// 反向遍历
	reverse(DoubleHead)

	traverse1(DoubleHead)

	// 查找已存在节点
	if lookUPDNode(DoubleHead, 5) {
		fmt.Println("该节点已存在!")
	} else {
		fmt.Println("该节点不存在!")
	}
}
