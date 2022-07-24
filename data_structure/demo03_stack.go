package main

import "fmt"

type StackNode struct {
	Value int
	Next  *StackNode
}

// 初始化栈结构（空栈）
var stackSize = 0
var stack = new(StackNode)

func Push(v int) bool {
	if stack == nil {
		stack = &StackNode{v, nil}
		stackSize = 1
		return true
	}

	temp := &StackNode{v, nil}
	temp.Next = stack
	stack = temp
	stackSize++
	return true
}

// 出栈
func Pop(t *StackNode) (int, bool) {

	if stackSize == 0 {
		return 0, false
	}

	//  只有一个节点
	if stackSize == 1 {
		stackSize = 0
		stack = nil
		return t.Value, true
	}

	// 将栈的头节点指针指向下一个节点，并返回之前的头节点数据
	stack = stack.Next
	stackSize--
	return t.Value, true
}

// 遍历（不删除任何节点，只读取值）
func traverse3(t *StackNode) {
	if stackSize == 0 {
		fmt.Println("空栈!")
		return
	}
	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func main() {
	stack = nil
	// 读取空栈
	v, b := Pop(stack)
	if b {
		fmt.Print(v, " ")
	} else {
		fmt.Println("Pop() 失败!")
	}
	// 进栈
	Push(100)
	// 遍历栈
	traverse3(stack)
	// 再次进栈
	Push(200)
	// 再次遍历栈
	traverse3(stack)
	// 批量进栈
	for i := 0; i < 10; i++ {
		Push(i)
	}
	// 批量出栈
	for i := 0; i < 15; i++ {
		v, b := Pop(stack)
		if b {
			fmt.Print(v, " ")
		} else {
			// 如果已经是空栈，则退出循环
			break
		}
	}
	fmt.Println()
	// 再次遍历栈
	traverse3(stack)
}
