package main

import "fmt"

func main() {
	ll := IntSLList{}
	ll.addToHead(6)
	ll.addToTail(20)
	ll.addToHead(6)
	ll.addToHead(6)
	ll.addToTail(5)
	ll.addToTail(20)
	ll.addToTail(5)
	ll.addToTail(8)
	ll.addToTail(5)
	ll.addToTail(20)
	ll.addToTail(5)

	ll.printAll()

	ll.removeDuplicates()

	fmt.Println("------------------------")

	ll.printAll()
}
