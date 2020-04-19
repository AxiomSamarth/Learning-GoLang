package main

import (
	"fmt"
)

type node struct {
	value int
	next  *node
}

// addLast function will add a new node to the linked list at the end
func (head *node) addLast(value int) *node {
	newNode := &node{value, nil}

	if head == nil {
		head = newNode
	} else {
		cursor := head
		for cursor.next != nil {
			cursor = cursor.next
		}
		cursor.next = newNode
	}
	return head
}

// addFirst function will add a new node at the beginning of the linked list
func (head *node) addFirst(value int) *node {
	newNode := &node{value, nil}

	if head == nil {
		head = newNode
	} else {
		newNode.next = head
		head = newNode
	}
	return head
}

// deleteFirst function deletes the node from the beginning of the linked list
func (head *node) deleteFirst() *node {
	if head == nil {
		fmt.Println("The linked list is empty")
		return nil
	}
	if head.next == nil {
		head = nil
	} else {
		head = head.next
	}
	return head
}

// deleteLast function deletes the last node from the linked list
// whose head is used to call the function
func (head *node) deleteLast() *node {
	if head == nil {
		fmt.Println("Linked list is empty")
		return nil
	}

	if head.next == nil {
		head = nil
	} else {
		cursor := head
		for cursor.next.next != nil {
			cursor = cursor.next
		}
		cursor.next = nil
	}
	return head
}

// reverseList prints the nodes of the linked list in reverse order
func (head *node) reverseList() *node {
	if head == nil {
		return head
	}
	head.next.reverseList()
	fmt.Printf("%v ", head.value)
	return head
}

// displayList will display the linked list values
func (head *node) displayList() {
	if head == nil {
		fmt.Println("The linked list is empty")
		return
	}
	fmt.Println()

	cursor := head
	for cursor != nil {
		fmt.Printf("%v ", cursor.value)
		cursor = cursor.next
	}
}

func main() {
	var head *node

	head = head.addLast(10)
	head = head.addLast(20)
	head = head.addLast(30)
	head = head.addFirst(40)

	head.displayList()

	head = head.deleteFirst()
	head = head.deleteLast()
	head.displayList()

	head = head.reverseList()
}
