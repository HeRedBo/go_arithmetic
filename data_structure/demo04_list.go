package main

import "fmt"

// 定义链表节点
type ListNode struct {
	Value int
	Next  *ListNode
}

// 初始化队列
var listSize = 0
var queue = new(ListNode)

// 入队 （从队头插入）
func ListPush(t *ListNode, v int) bool {
	if queue == nil {
		queue = &ListNode{v, nil}
		listSize++
		return true
	}

	t = &ListNode{v, nil}
	t.Next = queue
	queue = t
	listSize++
	return true
}

// 出队 （从队尾巴删除）
func ListPop(t *ListNode) (int, bool) {

	if listSize == 0 {
		return 0, false
	}

	if listSize == 1 {
		queue = nil
		listSize--
		return t.Value, true
	}

	// 迭代队列 直到队尾
	temp := t
	for (t.Next) != nil {
		temp = t
		t = t.Next
	}

	v := (temp.Next).Value
	temp.Next = nil
	listSize--
	return v, true
}

// 遍历队列所有节点
func listTraverse(t *ListNode) {

	if listSize == 0 {
		fmt.Println("空队列!")
		return
	}
	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func main() {
	queue = nil
	// 入队
	ListPush(queue, 10)
	fmt.Println("Size:", listSize)
	// 遍历
	listTraverse(queue)
	// 出队
	v, b := ListPop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", listSize)

	// 批量入队
	for i := 0; i < 5; i++ {
		ListPush(queue, i)
	}
	// 再次遍历
	listTraverse(queue)
	fmt.Println("Size:", listSize)
	// 出队
	v, b = ListPop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", listSize)
	// 再次出队
	v, b = ListPop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", listSize)
	// 再次遍历
	listTraverse(queue)
}
