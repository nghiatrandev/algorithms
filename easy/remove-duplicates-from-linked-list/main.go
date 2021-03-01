package main

import (
	"fmt"
)

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func main() {
	input := addMany(&LinkedList{Value: 1}, []int{1, 3, 4, 4, 4, 5, 6, 6})
	//expected := addMany(&LinkedList{Value: 1}, []int{3, 4, 5, 6})
	actual := RemoveDuplicatesFromLinkedList(input)
	fmt.Println(getValues(actual))
}

func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
	result := new(LinkedList)
	result.Value = linkedList.Value
	result.Next = linkedList.GetNext(linkedList.Value)
	return result
}

func (linkedList *LinkedList) GetNext(esceptValue int) *LinkedList {

	if linkedList != nil {
		//next := linkedList.GetNext(esceptValue)
		if linkedList.Value == esceptValue {
			linkedList.Next.GetNext(esceptValue)
		} else {
			return &LinkedList{
				Value: linkedList.Value,
				Next:  linkedList.GetNext(linkedList.Value),
			}
		}
	}

	return nil
}

// ############### sand box ############
func addMany(linkedList *LinkedList, values []int) *LinkedList {
	current := linkedList
	for current.Next != nil {
		current = current.Next
	}
	for _, value := range values {
		current.Next = &LinkedList{Value: value}
		current = current.Next
	}
	return linkedList
}

func getValues(linkedList *LinkedList) []int {
	values := []int{}
	current := linkedList
	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}
	return values
}
