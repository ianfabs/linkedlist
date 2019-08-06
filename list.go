package main

import "fmt"

type IntSLList struct {
	head *IntSLLNode
	tail *IntSLLNode
}

func (list *IntSLList) isEmpty() bool {
	return list.head == nil
}

func (list *IntSLList) addToHead(el int) {
	list.head = &IntSLLNode{info: el, next: list.head}
	if list.tail == nil {
		list.tail = list.head
	}
}

func (list *IntSLList) addToTail(el int) {
	if !list.isEmpty() {
		list.tail.next = &IntSLLNode{info: el}
		list.tail = list.tail.next
	} else {
		list.head = &IntSLLNode{info: el}
		list.tail = &IntSLLNode{info: el}
	}
}

func (list *IntSLList) deleteFromHead() int {
	el := list.head.info
	if list.head == list.tail {
		list.head = nil
		list.tail = nil
	} else {
		list.head = list.head.next
	}
	return el
}

func (list *IntSLList) deleteFromTail() int {
	el := list.tail.info
	if list.head == list.tail {
		list.head = nil
		list.tail = nil
	} else {
		var tmp *IntSLLNode
		for tmp = list.head; tmp.next != list.tail; tmp = tmp.next {
		}
		list.tail = tmp
		list.tail.next = nil
	}
	return el
}
func (list *IntSLList) printAll() {
	for tmp := list.head; tmp != nil; tmp = tmp.next {
		fmt.Println(tmp.info)
	}
}

func (list *IntSLList) isInList(el int) bool {
	var tmp *IntSLLNode
	for tmp = list.head; tmp != nil && tmp.info != el; tmp = tmp.next {
	}
	return tmp != nil
}
func (list *IntSLList) delete(el int) {
	if !list.isEmpty() {
		if list.head == list.tail && el == list.head.info {
			list.head = nil
			list.tail = nil
		} else if el == list.head.info {
			list.head = list.head.next
		} else {
			//var pred, tmp *IntSLLNode
			pred, tmp := list.head, list.head.next
			for tmp != nil && tmp.info != el {
				pred, tmp = pred.next, tmp.next
			}
			if tmp != nil {
				pred.next = tmp.next
				if tmp == list.tail {
					list.tail = pred
				}
			}

		}
	}
}
func (list *IntSLList) sum() int {
	var sum int
	for tmp := list.head; tmp != nil; tmp = tmp.next {
		sum += tmp.info
	}
	return sum
}
func (list *IntSLList) max() int {
	max := list.head.info
	for tmp := list.head; tmp != nil; tmp = tmp.next {
		if max < tmp.info {
			max = tmp.info
		}
	}
	return max
}
func (list *IntSLList) min() int {
	max := list.head.info
	for tmp := list.head; tmp != nil; tmp = tmp.next {
		if max > tmp.info {
			max = tmp.info
		}
	}
	return max
}

func (list *IntSLList) removeDuplicates() {
	var m = make(map[int]bool)

	var pn *IntSLLNode
	cn := list.head
	for cn != nil {
		if m[cn.info] {
			pn.next = cn.next
		} else {
			m[cn.info] = true
			pn = cn
		}

		cn = cn.next
	}
}
