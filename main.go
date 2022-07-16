package main

import (
	"fmt"
	"go_pizza/linked_list"
)

func main(){
	fmt.Println("hello pizza")
	NewList := linked_list.NewLinkedList()
	NewList.Enqueue("cheese")
	NewList.Enqueue("pepperoni")
	NewList.Enqueue("stuff")
	fmt.Println(NewList.Length)
	/*
	popped := NewList.Dequeue()
	fmt.Println(NewList.Length)
	fmt.Println(popped.Name)
	popped2 := NewList.Dequeue()
	fmt.Println(NewList.Length)
	fmt.Println(popped2.Name)
	popped3 := NewList.Dequeue()
	fmt.Println(NewList.Length)
	fmt.Println(popped3.Name)
	*/
	deleted := NewList.Delete("cheese")
	fmt.Println(deleted)
	fmt.Println(NewList.Length)
}