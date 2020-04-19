package main

import "fmt"

type node struct {
	value int
	next  *node
}

func (top *node) push(value int) *node {
	newNode := &node{value, nil}
	newNode.next = top
	top = newNode
	return top
}

func (top *node) pop() *node {
	if top == nil {
		fmt.Println("The stack is empty")
		return nil
	}
	fmt.Printf("\nPopping %v from the stack", top.value)
	top = top.next
	return top
}

func (top *node) showPeak() {
	if top == nil {
		fmt.Println("The stack is empty to show the peak")
		return
	}
	fmt.Printf("\nThe peak of the stack is %v", top.value)
}

func (top *node) displayStack() {
	if top == nil {
		fmt.Println("Stack is empty")
		return
	}
	cursor := top
	for cursor != nil {
		fmt.Printf("%v ", cursor.value)
		cursor = cursor.next
	}
}

func main() {
	var top *node

	top = top.push(10)
	top = top.push(20)
	top = top.push(30)

	top.displayStack()

	top = top.pop()
	top.showPeak()
}
